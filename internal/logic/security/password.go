package security

import (
	"context"

	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/util/encrypt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// 全局实例
var (
	PwdManager = &passwordManager{}
)

type passwordManager struct {
}

// PasswordPolicy 密码策略
type PasswordPolicy struct {
	MinLength     int  `json:"minLength"`     // 最小长度
	RequireUpper  bool `json:"requireUpper"`  // 需要大写字母
	RequireLower  bool `json:"requireLower"`  // 需要小写字母
	RequireDigit  bool `json:"requireDigit"`  // 需要数字
	RequireSymbol bool `json:"requireSymbol"` // 需要特殊字符
	MaxAgeDays    int  `json:"maxAgeDays"`    // 密码最大有效期
	HistoryCount  int  `json:"historyCount"`  // 历史密码记录数量
}

// CreatePassword 创建新密码（注册时使用）
func (p *passwordManager) CreatePassword(ctx context.Context, memberID int64, plainPassword string) error {
	// 1. 验证密码强度
	policy, err := p.GetPasswordPolicy(ctx)
	if err != nil {
		return err
	}

	strength := encrypt.PasswordStrength{
		MinLength:     policy.MinLength,
		RequireUpper:  policy.RequireUpper,
		RequireLower:  policy.RequireLower,
		RequireDigit:  policy.RequireDigit,
		RequireSymbol: policy.RequireSymbol,
	}

	if err := encrypt.ValidatePasswordStrength(plainPassword, strength); err != nil {
		return err
	}

	// 2. 生成bcrypt哈希
	hash, err := encrypt.HashPassword(plainPassword)
	if err != nil {
		return err
	}

	// 3. 更新用户密码
	_, err = dao.MemberInfo.Ctx(ctx).Data(gdb.Map{
		"password_hash": hash,
		"salt":          "", // bcrypt不需要单独的盐值
		"updated_at":    gtime.Now(),
	}).WherePri(memberID).Update()

	if err != nil {
		return gerror.Wrap(err, "更新密码失败")
	}

	return nil
}

// ChangePassword 修改密码
func (p *passwordManager) ChangePassword(ctx context.Context, memberID int64, oldPassword, newPassword string) error {
	// 1. 验证旧密码
	member, err := dao.MemberInfo.FindByID(ctx, memberID)
	if err != nil {
		return gerror.Wrap(err, "获取用户信息失败")
	}

	if member == nil {
		return gerror.New("用户不存在")
	}

	// 验证旧密码
	if err := p.verifyCurrentPassword(ctx, member, oldPassword); err != nil {
		return err
	}

	// 2. 验证新密码强度
	policy, err := p.GetPasswordPolicy(ctx)
	if err != nil {
		return err
	}

	strength := encrypt.PasswordStrength{
		MinLength:     policy.MinLength,
		RequireUpper:  policy.RequireUpper,
		RequireLower:  policy.RequireLower,
		RequireDigit:  policy.RequireDigit,
		RequireSymbol: policy.RequireSymbol,
	}

	if err := encrypt.ValidatePasswordStrength(newPassword, strength); err != nil {
		return err
	}

	// 3. 生成新的bcrypt哈希
	newHash, err := encrypt.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 4. 更新密码
	_, err = dao.MemberInfo.Ctx(ctx).Data(gdb.Map{
		"password_hash": newHash,
		"salt":          "", // 清空旧的盐值
		"updated_at":    gtime.Now(),
	}).WherePri(memberID).Update()

	if err != nil {
		return gerror.Wrap(err, "更新密码失败")
	}

	return nil
}

// ResetPasswordRequest 申请密码重置
func (p *passwordManager) ResetPasswordRequest(ctx context.Context, email string) (string, error) {
	// 1. 查找用户
	member, err := dao.MemberInfo.Ctx(ctx).Where("email", email).One()
	if err != nil {
		return "", gerror.Wrap(err, "查询用户失败")
	}

	if member.IsEmpty() {
		return "", gerror.New("邮箱不存在")
	}

	// 2. 生成重置令牌
	token, err := encrypt.GenerateSecureToken(32)
	if err != nil {
		return "", err
	}

	// 3. 保存重置令牌
	_, err = dao.MemberInfo.Ctx(ctx).Data(gdb.Map{
		"password_reset_token": token,
		"updated_at":           gtime.Now(),
	}).Where("email", email).Update()

	if err != nil {
		return "", gerror.Wrap(err, "保存重置令牌失败")
	}

	return token, nil
}

// ResetPasswordConfirm 确认密码重置
func (p *passwordManager) ResetPasswordConfirm(ctx context.Context, token, newPassword string) error {
	// 1. 验证重置令牌
	member, err := dao.MemberInfo.Ctx(ctx).Where("password_reset_token", token).One()
	if err != nil {
		return gerror.Wrap(err, "查询重置令牌失败")
	}

	if member.IsEmpty() {
		return gerror.New("重置令牌无效")
	}

	// 2. 验证新密码
	policy, err := p.GetPasswordPolicy(ctx)
	if err != nil {
		return err
	}

	strength := encrypt.PasswordStrength{
		MinLength:     policy.MinLength,
		RequireUpper:  policy.RequireUpper,
		RequireLower:  policy.RequireLower,
		RequireDigit:  policy.RequireDigit,
		RequireSymbol: policy.RequireSymbol,
	}

	if err := encrypt.ValidatePasswordStrength(newPassword, strength); err != nil {
		return err
	}

	// 3. 生成新密码哈希
	newHash, err := encrypt.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 4. 获取用户ID
	var memberEntity *entity.MemberInfo
	if err := member.Struct(&memberEntity); err != nil {
		return gerror.Wrap(err, "数据转换失败")
	}

	// 5. 更新密码并清除重置令牌
	_, err = dao.MemberInfo.Ctx(ctx).Data(gdb.Map{
		"password_hash":        newHash,
		"salt":                 "",
		"password_reset_token": "",
		"updated_at":           gtime.Now(),
	}).WherePri(memberEntity.Id).Update()

	if err != nil {
		return gerror.Wrap(err, "重置密码失败")
	}

	return nil
}

// GetPasswordPolicy 获取密码策略
func (p *passwordManager) GetPasswordPolicy(ctx context.Context) (*PasswordPolicy, error) {
	// 从数据库获取密码策略配置
	// 这里简化处理，实际应该从配置表获取
	return &PasswordPolicy{
		MinLength:     8,
		RequireUpper:  true,
		RequireLower:  true,
		RequireDigit:  true,
		RequireSymbol: false,
		MaxAgeDays:    90,
		HistoryCount:  5,
	}, nil
}

// verifyCurrentPassword 验证当前密码
func (p *passwordManager) verifyCurrentPassword(ctx context.Context, member *entity.MemberInfo, password string) error {
	// 如果是bcrypt哈希
	if len(member.PasswordHash) > 50 && (member.PasswordHash[:4] == "$2a$" || member.PasswordHash[:4] == "$2b$" || member.PasswordHash[:4] == "$2y$") {
		return encrypt.VerifyPassword(password, member.PasswordHash)
	}

	// 兼容旧的MD5方式
	return gerror.New("请联系管理员升级您的账户安全")
}
