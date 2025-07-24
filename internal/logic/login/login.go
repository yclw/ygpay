package login

import (
	"context"
	"time"
	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/internal/consts"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/global"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/pkg/token"
	"yclw/ygpay/util/encrypt"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

var LoginService = NewLogin()

type Login struct {
}

func NewLogin() *Login {
	return &Login{}
}

func (l *Login) AccountLogin(ctx context.Context, username, password string) (res *v1.AccountLoginRes, err error) {
	res = &v1.AccountLoginRes{}
	// 获取用户信息
	var mb *entity.MemberInfo
	if mb, err = dao.MemberInfo.FindByUsername(ctx, username); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败")
		return
	}
	if mb == nil {
		err = gerror.New("用户名或密码错误")
		return
	}

	// 验证密码
	if err = checkPassword(password, mb.PasswordHash); err != nil {
		err = gerror.New("用户名或密码错误")
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用")
		return
	}
	res.LoginRes, err = l.handleLogin(ctx, mb)
	return
}

// checkPassword 检查密码
func checkPassword(input, hash string) (err error) {
	// 解码密码
	password, err := gbase64.Decode([]byte(input))
	if err != nil {
		err = gerror.New("密码解析失败")
		return
	}
	// 解密密码
	password, err = encrypt.AesECBDecrypt(password, []byte(consts.RequestEncryptKey))
	if err != nil {
		err = gerror.New("密码解析失败")
		return
	}

	plainPassword := string(password)

	// bcrypt验证方式
	if isBcryptHash(hash) {
		err = encrypt.VerifyPassword(plainPassword, hash)
		return
	}
	err = gerror.New("用户名或密码错误")
	return
}

// isBcryptHash 判断是否为bcrypt哈希
func isBcryptHash(hash string) bool {
	// bcrypt哈希格式：$2a$12$...（通常60-70字符）
	return len(hash) > 50 && (hash[:4] == "$2a$" || hash[:4] == "$2b$" || hash[:4] == "$2y$")
}

func (l *Login) handleLogin(ctx context.Context, member *entity.MemberInfo) (res *v1.LoginRes, err error) {
	// 获取用户角色
	roleId, err := dao.MemberRole.FindRoleIdByMemberId(ctx, member.Id)
	if err != nil {
		err = gerror.Wrap(err, "获取用户角色失败")
		return
	}
	// 获取角色信息
	role, err := dao.RoleInfo.FindByID(ctx, roleId)
	if err != nil {
		err = gerror.Wrap(err, "获取角色信息失败")
		return
	}

	// 创建用户身份
	user := &token.AccessIdentity{
		Uid:      member.Uid,
		RoleId:   role.Id,
		RoleKey:  role.Key,
		Username: member.Username,
		Avatar:   member.Avatar,
		Email:    member.Email,
		Mobile:   member.Mobile,
		App:      global.AppName(ctx),
		LoginAt:  gtime.Now(),
	}

	// 生成访问token
	header, expires, err := token.AccessJwt.CreateToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	// 生成刷新token
	refreshUser := &token.RefreshIdentity{
		Uid: user.Uid,
	}
	refreshHeader, refreshExpires, err := token.RefreshJwt.CreateToken(ctx, *refreshUser)
	if err != nil {
		return nil, err
	}
	refreshKey := token.RefreshJwt.GetCacheKey(ctx, user.Uid)
	global.Cache().Set(ctx, refreshKey, refreshHeader, time.Duration(refreshExpires)*time.Second)

	// 返回登录信息
	res = &v1.LoginRes{
		Username:     user.Username,
		Uid:          user.Uid,
		Token:        header,
		Expires:      expires,
		RefreshToken: refreshHeader,
	}
	return
}

// Logout 登出
func (l *Login) Logout(ctx context.Context, header string) (err error) {
	// 删除访问token缓存
	ok, identity, err := token.AccessJwt.VerifyToken(ctx, header)
	if err != nil || !ok {
		return
	}
	// 删除刷新token缓存
	refreshKey := token.RefreshJwt.GetCacheKey(ctx, identity.Uid)
	global.Cache().Remove(ctx, refreshKey)
	return
}

// RefreshToken 刷新token
func (l *Login) RefreshToken(ctx context.Context, refreshToken string) (res *v1.LoginRefreshTokenRes, err error) {
	// 验证刷新token
	ok, identity, err := token.RefreshJwt.VerifyToken(ctx, refreshToken)
	if err != nil || !ok {
		return
	}
	// 验证刷新token是否在缓存中
	refreshKey := token.RefreshJwt.GetCacheKey(ctx, identity.Uid)
	cacheToken, err := global.Cache().Get(ctx, refreshKey)
	if err != nil {
		return
	}
	if cacheToken.String() != refreshToken {
		return nil, gerror.New("token已失效，请重新登录")
	}

	// 获取用户角色
	member, err := dao.MemberInfo.FindByUid(ctx, identity.Uid)
	if err != nil {
		err = gerror.Wrap(err, "获取用户信息失败")
		return
	}
	// 获取角色信息
	roleId, err := dao.MemberRole.FindRoleIdByMemberId(ctx, member.Id)
	if err != nil {
		err = gerror.Wrap(err, "获取用户角色失败")
		return
	}
	role, err := dao.RoleInfo.FindByID(ctx, roleId)
	if err != nil {
		err = gerror.Wrap(err, "获取角色信息失败")
		return
	}

	// 创建用户身份
	user := &token.AccessIdentity{
		Uid:      member.Uid,
		RoleId:   role.Id,
		RoleKey:  role.Key,
		Username: member.Username,
		Avatar:   member.Avatar,
		Email:    member.Email,
		Mobile:   member.Mobile,
		App:      global.AppName(ctx),
		LoginAt:  gtime.Now(),
	}

	// 生成访问token
	header, expires, err := token.AccessJwt.CreateToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	// 生成刷新token
	refreshUser := &token.RefreshIdentity{
		Uid: user.Uid,
	}
	refreshHeader, refreshExpires, err := token.RefreshJwt.CreateToken(ctx, *refreshUser)
	if err != nil {
		return nil, err
	}
	// 更新刷新token缓存
	global.Cache().Set(ctx, refreshKey, refreshHeader, time.Duration(refreshExpires)*time.Second)

	// 返回登录信息
	res = &v1.LoginRefreshTokenRes{
		Token:        header,
		Expires:      expires,
		RefreshToken: refreshHeader,
	}
	return
}
