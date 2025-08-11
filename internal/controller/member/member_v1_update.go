package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/internal/logic/member"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	model, err := c.updateReqToUpdateModel(req)
	if err != nil {
		return
	}

	// 根据roleUid获取roleId
	roleId, err := c.RoleService.GetRoleIdByUid(ctx, req.RoleUid)
	if err != nil {
		return
	}

	// 设置角色ID
	model.MemberRole.RoleId = roleId

	err = c.MemberService.Update(ctx, model)
	if err != nil {
		return
	}

	res = &v1.UpdateRes{}
	return
}

// updateReqToUpdateModel 将更新请求转换为更新模型
func (c *ControllerV1) updateReqToUpdateModel(req *v1.UpdateReq) (*member.MemberUpdateModel, error) {
	return &member.MemberUpdateModel{
		Password: req.Password,
		Uid:      req.Uid,
		MemberInfo: &do.MemberInfo{
			Uid:      req.Uid,
			Username: req.Username,
			Nickname: req.Nickname,
			Avatar:   req.Avatar,
			Sex:      req.Sex,
			Email:    req.Email,
			Mobile:   req.Mobile,
			Address:  req.Address,
			Remark:   req.Remark,
			Sort:     req.Sort,
			Status:   req.Status,
		},
	}, nil
}
