package user

import (
	"context"
	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) UserInfo(ctx context.Context, _ *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	memberId := contexts.GetUserId(ctx)
	member, err := c.MemberService.GetOneEncrypt(ctx, memberId)
	if err != nil {
		return
	}
	res = &v1.UserInfoRes{
		Id:          member.Id,
		RoleName:    member.RoleName,
		Permissions: member.Permissions,
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
	return
}
