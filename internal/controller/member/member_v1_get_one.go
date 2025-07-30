package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	uid := contexts.GetUserUid(ctx)
	member, err := c.MemberService.GetOne(ctx, uid)
	if err != nil {
		return
	}
	res = &v1.GetOneRes{}
	res.Member = &v1.MemberModel{
		Uid:          member.Uid,
		RoleId:       member.RoleId,
		Username:     member.Username,
		PasswordHash: member.PasswordHash,
		Nickname:     member.Nickname,
		Avatar:       member.Avatar,
		Sex:          member.Sex,
		Email:        member.Email,
		Mobile:       member.Mobile,
		Address:      member.Address,
		LastActiveAt: member.LastActiveAt,
		Remark:       member.Remark,
		Sort:         member.Sort,
		Status:       member.Status,
		CreatedAt:    member.CreatedAt,
		UpdatedAt:    member.UpdatedAt,
		LoginCount:   member.LoginCount,
		LastLoginAt:  member.LastLoginAt,
		LastLoginIp:  member.LastLoginIp,
	}
	return
}
