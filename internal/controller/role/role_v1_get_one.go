package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/role"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	model, err := c.RoleService.GetOne(ctx, req.Id)
	if err != nil {
		return
	}
	res = &v1.GetOneRes{
		RoleModel: c.roleModelToV1(model),
	}
	return
}

func (c *ControllerV1) roleModelToV1(model *role.RoleModel) *v1.RoleModel {
	return &v1.RoleModel{
		Id:        model.RoleInfo.Id,
		Key:       model.RoleInfo.Key,
		Name:      model.RoleInfo.Name,
		Remark:    model.RoleInfo.Remark,
		Sort:      model.RoleInfo.Sort,
		Status:    model.RoleInfo.Status,
		CreatedAt: model.RoleInfo.CreatedAt,
		UpdatedAt: model.RoleInfo.UpdatedAt,
	}
}
