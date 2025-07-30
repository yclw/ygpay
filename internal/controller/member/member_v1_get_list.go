package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	members, err := c.MemberService.GetAllList(ctx)
	if err != nil {
		return
	}
	res = &v1.GetListRes{}
	res.List = make([]*v1.MemberModel, 0, len(members))
	for _, member := range members {
		res.List = append(res.List, &v1.MemberModel{
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
		})
	}
	return
}
