package login

import (
	"context"
	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/internal/consts"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/global"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/pkg/token"
	"yclw/ygpay/util/simple"

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

	if mb.Salt == "" {
		err = gerror.New("用户信息错误")
		return
	}

	if err = simple.CheckPassword(password, mb.Salt, mb.PasswordHash); err != nil {
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

func (l *Login) handleLogin(ctx context.Context, member *entity.MemberInfo) (res *v1.LoginRes, err error) {
	role, err := dao.RoleInfo.FindByID(ctx, member.RoleId)
	if err != nil {
		err = gerror.Wrap(err, "获取角色信息失败")
		return
	}

	// 创建用户身份
	user := &token.Identity{
		Id:       member.Id,
		RoleId:   role.Id,
		RoleKey:  role.Key,
		Username: member.Username,
		Avatar:   member.Avatar,
		Email:    member.Email,
		Mobile:   member.Mobile,
		App:      global.AppName(ctx),
		LoginAt:  gtime.Now(),
	}

	// 生成登录token
	header, expires, err := token.Jwt.CreateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	// TODO: 添加到token缓存

	// 返回登录信息
	res = &v1.LoginRes{
		Username: user.Username,
		Id:       user.Id,
		Token:    header,
		Expires:  expires,
	}
	return
}

// Logout 登出
func (l *Login) Logout(ctx context.Context, header string) (err error) {

	// TODO: 删除token缓存

	return
}
