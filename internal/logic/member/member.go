package member

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var MemberService = NewMember()

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

// MemberModel 用户模型
type MemberModel struct {
	Uid         string      `json:"uid"                dc:"用户ID"`
	RoleId      int64       `json:"roleId"             dc:"所属角色"`
	Permissions []string    `json:"permissions"        dc:"权限信息"`
	Username    string      `json:"username"           dc:"用户名"`
	Avatar      string      `json:"avatar"             dc:"头像"`
	Sex         int         `json:"sex"                dc:"性别"`
	Email       string      `json:"email"              dc:"邮箱"`
	Mobile      string      `json:"mobile"             dc:"手机号码"`
	Address     string      `json:"address"            dc:"联系地址"`
	CreatedAt   *gtime.Time `json:"createdAt"          dc:"创建时间"`
	*MemberLoginStatModel
}

// MemberLoginStat 用户登录统计
type MemberLoginStatModel struct {
	LoginCount  int         `json:"loginCount"  dc:"登录次数"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
	LastLoginIp string      `json:"lastLoginIp" dc:"最后登录IP"`
}

// CreateMemberModel 创建用户模型
type CreateMemberModel struct {
	Username     string `json:"username"           dc:"用户名"`
	PasswordHash string `json:"passwordHash"       dc:"密码"`
	RoleId       int64  `json:"roleId"             dc:"角色ID"`
	Avatar       string `json:"avatar"             dc:"头像"`
	Sex          int    `json:"sex"                dc:"性别"`
	Email        string `json:"email"              dc:"邮箱"`
	Mobile       string `json:"mobile"             dc:"手机号码"`
	Address      string `json:"address"            dc:"联系地址"`
	Remark       string `json:"remark"             dc:"备注"`
	Status       int    `json:"status"             dc:"状态"`
}

// UpdateMemberModel 更新用户模型
type UpdateMemberModel struct {
	Uid          string `json:"uid"                dc:"用户ID"`
	Username     string `json:"username"           dc:"用户名"`
	PasswordHash string `json:"passwordHash"       dc:"密码"`
	RoleId       int64  `json:"roleId"             dc:"角色ID"`
	Avatar       string `json:"avatar"             dc:"头像"`
	Sex          int    `json:"sex"                dc:"性别"`
	Email        string `json:"email"              dc:"邮箱"`
	Mobile       string `json:"mobile"             dc:"手机号码"`
	Address      string `json:"address"            dc:"联系地址"`
	Remark       string `json:"remark"             dc:"备注"`
	Status       int    `json:"status"             dc:"状态"`
}

// GetOne 获取单个用户信息
func (m *Member) GetOne(ctx context.Context, uid string) (res *MemberModel, err error) {
	res = &MemberModel{}

	if uid == "" {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	member, err := dao.MemberInfo.FindByUid(ctx, uid)
	if err != nil {
		err = gerror.Wrap(err, "获取用户信息失败")
		return
	}
	if member == nil {
		err = gerror.New("用户不存在！")
		return
	}

	// 使用 gconv 复制字段
	if err = gconv.Struct(member, res); err != nil {
		err = gerror.Wrap(err, "字段复制失败")
		return
	}

	// 获取登录信息
	stat, err := m.getLoginStat(ctx, member.Id)
	if err != nil {
		return
	}
	res.LoginCount = stat.LoginCount
	res.LastLoginAt = stat.LastLoginAt
	res.LastLoginIp = stat.LastLoginIp

	return
}

// GetOneEncrypt 获取单个用户信息--加密处理
func (m *Member) GetOneEncrypt(ctx context.Context, uid string) (res *MemberModel, err error) {
	res, err = m.GetOne(ctx, uid)
	if err != nil {
		return
	}
	res.Mobile = gstr.HideStr(res.Mobile, 40, `*`) // 手机号脱敏
	res.Email = gstr.HideStr(res.Email, 40, `*`)   // 邮箱脱敏
	return
}

// getLoginStat 获取用户登录统计
func (s *Member) getLoginStat(ctx context.Context, memberId int64) (res *MemberLoginStatModel, err error) {
	last, err := dao.LogLogin.FindLastByMemberId(ctx, memberId)
	if err != nil {
		return
	}
	res = &MemberLoginStatModel{}
	res.LastLoginAt = last.LoginTime
	res.LastLoginIp = last.IpAddress
	res.LoginCount, err = dao.LogLogin.GetLoginCount(ctx, memberId)
	return
}

// GetAllList 获取所有用户列表
func (s *Member) GetAllList(ctx context.Context) (res []*MemberModel, err error) {
	members, err := dao.MemberInfo.FindAll(ctx)
	if err != nil {
		return
	}

	res = make([]*MemberModel, 0, len(members))
	for _, member := range members {
		memberModel := &MemberModel{}

		// 使用 gconv 复制字段
		if err = gconv.Struct(member, memberModel); err != nil {
			err = gerror.Wrap(err, "字段复制失败")
			return
		}

		// 获取登录统计信息
		stat, err := s.getLoginStat(ctx, member.Id)
		if err != nil {
			return nil, err
		}
		memberModel.LoginCount = stat.LoginCount
		memberModel.LastLoginAt = stat.LastLoginAt
		memberModel.LastLoginIp = stat.LastLoginIp

		res = append(res, memberModel)
	}
	return
}

// GetSubList 获取子用户列表
func (s *Member) GetSubList(ctx context.Context, memberId int64) (res []*MemberModel, err error) {
	// 获取用户角色
	roleId, err := dao.MemberRole.FindRoleIdByMemberId(ctx, memberId)
	if err != nil {
		return
	}
	// 获取子角色id列表
	subRoleIds, err := dao.RoleTree.FindSubRoleIds(ctx, roleId)
	if err != nil {
		return
	}
	// 获取用户ID列表
	userIds, err := dao.MemberRole.FindUserIdsByRoleIds(ctx, subRoleIds)
	if err != nil {
		return
	}

	// 获取子用户列表
	members, err := dao.MemberInfo.FindByUserIds(ctx, userIds)
	if err != nil {
		return
	}

	res = make([]*MemberModel, 0, len(members))
	for _, member := range members {
		memberModel := &MemberModel{}

		// 使用 gconv 复制字段
		if err = gconv.Struct(member, memberModel); err != nil {
			err = gerror.Wrap(err, "字段复制失败")
			return
		}

		// 获取登录统计信息
		stat, err := s.getLoginStat(ctx, member.Id)
		if err != nil {
			return nil, err
		}
		memberModel.LoginCount = stat.LoginCount
		memberModel.LastLoginAt = stat.LastLoginAt
		memberModel.LastLoginIp = stat.LastLoginIp

		res = append(res, memberModel)
	}
	return
}

// Create 创建用户
func (s *Member) Create(ctx context.Context, req *CreateMemberModel) (err error) {
	member := &entity.MemberInfo{}
	if err = gconv.Struct(req, member); err != nil {
		err = gerror.Wrap(err, "字段复制失败")
		return
	}

	// 创建用户
	if err = dao.MemberInfo.Create(ctx, member); err != nil {
		err = gerror.Wrap(err, "创建用户失败")
		return
	}
	return
}

// Update 更新用户
func (s *Member) Update(ctx context.Context, req *UpdateMemberModel) (err error) {
	return
}

// Delete 删除用户
func (s *Member) Delete(ctx context.Context, memberId int64) (err error) {
	return
}
