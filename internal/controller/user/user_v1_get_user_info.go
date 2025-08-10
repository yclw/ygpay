package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	uid := contexts.GetUserUid(ctx)

	// 获取用户信息
	member, err := c.MemberService.GetOneEncrypt(ctx, uid)
	if err != nil {
		return
	}

	// 获取角色信息
	role, err := c.RoleService.GetOne(ctx, member.RoleId)
	if err != nil {
		return
	}

	res = &v1.GetUserInfoRes{
		Uid:       uid,
		Nickname:  member.Nickname,
		RoleName:  role.Name,
		Username:  member.Username,
		Avatar:    member.Avatar,
		Sex:       member.Sex,
		Email:     member.Email,
		Mobile:    member.Mobile,
		Address:   member.Address,
		CreatedAt: member.MemberInfo.CreatedAt,
	}
	return
}
