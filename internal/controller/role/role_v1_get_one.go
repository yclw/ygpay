package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/role"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	// 获取角色信息
	model, err := c.RoleService.GetOne(ctx, req.Id)
	if err != nil {
		return
	}

	// 转换为v1.RoleModel
	res = &v1.GetOneRes{
		RoleModel: c.roleModelToV1(model),
	}
	return
}

// roleModelToV1 将role.RoleModel转换为v1.RoleModel
func (c *ControllerV1) roleModelToV1(model *role.RoleModel) *v1.RoleModel {
	return &v1.RoleModel{
		Id:         model.RoleInfo.Id,
		Name:       model.RoleInfo.Name,
		Key:        model.RoleInfo.Key,
		ParentId:   model.RoleInfo.Pid,
		ParentName: model.ParentName,
		Remark:     model.RoleInfo.Remark,
		Sort:       model.RoleInfo.Sort,
		Status:     model.RoleInfo.Status,
		CreatedAt:  model.RoleInfo.CreatedAt,
		UpdatedAt:  model.RoleInfo.UpdatedAt,
	}
}
