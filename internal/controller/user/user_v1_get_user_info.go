package user

import (
	"context"

	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	uid := contexts.GetUserUid(ctx)
	member, err := c.MemberService.GetOneEncrypt(ctx, uid)
	if err != nil {
		return
	}
	res = &v1.GetUserInfoRes{
		Uid: uid,
		// RoleId:      member.RoleId,
		// Permissions: member.Permissions,
		Username:    member.Username,
		Avatar:      member.Avatar,
		Sex:         member.Sex,
		Email:       member.Email,
		Mobile:      member.Mobile,
		Address:     member.Address,
		CreatedAt:   member.CreatedAt,
		LoginCount:  member.LoginCount,
		LastLoginAt: member.LastLoginAt,
		LastLoginIp: member.LastLoginIp,
	}
	res.RoleId = contexts.GetRoleId(ctx)
	return
}
