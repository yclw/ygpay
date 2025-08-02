package user

import (
	"context"

	v2 "yclw/ygpay/api/user/v2"
	"yclw/ygpay/pkg/contexts"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV2) GetUserInfo(ctx context.Context, req *v2.GetUserInfoReq) (res *v2.GetUserInfoRes, err error) {
	uid := contexts.GetUserUid(ctx)

	// 获取用户信息
	member, err := c.MemberService.GetOneEncrypt(ctx, uid)
	if err != nil {
		return
	}

	if member.RoleId != contexts.GetRoleId(ctx) {
		return nil, gerror.New("用户角色不匹配,请重新登录")
	}

	// // 获取角色信息
	role, err := c.RoleService.GetOne(ctx, member.RoleId)
	if err != nil {
		return
	}

	res = &v2.GetUserInfoRes{
		Uid:       uid,
		Nickname:  member.Nickname,
		RoleName:  role.Name,
		Username:  member.Username,
		Avatar:    member.Avatar,
		Sex:       member.Sex,
		Email:     member.Email,
		Mobile:    member.Mobile,
		Address:   member.Address,
		CreatedAt: member.CreatedAt,
	}
	return
}
