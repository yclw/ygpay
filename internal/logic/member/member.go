package member

import (
	"context"
	"yclw/ygpay/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

var MemberService = NewMember()

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

// MemberModel 用户模型
type MemberModel struct {
	Id          int64       `json:"id"                 dc:"用户ID"`
	RoleName    string      `json:"roleName"           dc:"所属角色"`
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

// GetOne 获取单个用户信息
func (m *Member) GetOne(ctx context.Context, memberId int64) (res *MemberModel, err error) {
	res = &MemberModel{}

	if memberId <= 0 {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	member, err := dao.MemberInfo.FindByID(ctx, memberId)
	if err != nil {
		err = gerror.Wrap(err, "获取用户信息失败")
		return
	}
	if member == nil {
		err = gerror.New("用户不存在！")
		return
	}

	// 获取登录信息
	stat, err := m.loginStat(ctx, memberId)
	if err != nil {
		return
	}
	res.LoginCount = stat.LoginCount
	res.LastLoginAt = stat.LastLoginAt
	res.LastLoginIp = stat.LastLoginIp

	return
}

// GetOneEncrypt 获取单个用户信息--加密处理
func (m *Member) GetOneEncrypt(ctx context.Context, memberId int64) (res *MemberModel, err error) {
	res, err = m.GetOne(ctx, memberId)
	if err != nil {
		return
	}
	res.Mobile = gstr.HideStr(res.Mobile, 40, `*`) // 手机号脱敏
	res.Email = gstr.HideStr(res.Email, 40, `*`)   // 邮箱脱敏
	return
}

// loginStat 用户登录统计
func (s *Member) loginStat(ctx context.Context, memberId int64) (res *MemberLoginStatModel, err error) {
	last, err := dao.LogLogin.FindLastByMemberId(ctx, memberId)
	if err != nil {
		return
	}
	res = &MemberLoginStatModel{}
	res.LastLoginAt = last.LoginAt
	res.LastLoginIp = last.LoginIp
	res.LoginCount, err = dao.LogLogin.GetLoginCount(ctx, memberId)
	return
}
