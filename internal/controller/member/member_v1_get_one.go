package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/internal/logic/member"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	member, err := c.MemberService.GetOne(ctx, req.Uid)
	if err != nil {
		return
	}

	// 转换基本用户信息
	memberModel := c.memberModelToV1(member)

	// 获取角色名称
	role, roleErr := c.RoleService.GetOne(ctx, member.RoleId)
	if roleErr == nil && role != nil && role.RoleInfo != nil {
		memberModel.RoleName = role.RoleInfo.Name
	}

	res = &v1.GetOneRes{
		MemberModel: memberModel,
	}
	return
}

func (c *ControllerV1) memberModelToV1(member *member.MemberModel) *v1.MemberModel {
	return &v1.MemberModel{
		Uid:          member.Uid,
		RoleId:       member.RoleId,
		Username:     member.Username,
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
	}
}
