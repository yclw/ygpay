package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/internal/logic/member"
	"yclw/ygpay/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	// 将创建请求转换为创建模型
	model, err := c.createReqToCreateModel(req)
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

	// 创建用户
	err = c.MemberService.Create(ctx, model)
	if err != nil {
		return
	}

	res = &v1.CreateRes{}
	return
}

// createReqToCreateModel 将创建请求转换为创建模型
func (c *ControllerV1) createReqToCreateModel(req *v1.CreateReq) (*member.MemberCreateModel, error) {
	return &member.MemberCreateModel{
		MemberInfo: &do.MemberInfo{
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
