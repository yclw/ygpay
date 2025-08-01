package role

import (
	"context"

	v2 "yclw/ygpay/api/role/v2"
	"yclw/ygpay/internal/logic/role"
)

func (c *ControllerV2) GetList(ctx context.Context, req *v2.GetListReq) (res *v2.GetListRes, err error) {
	// 构建筛选参数
	filter := &role.RoleListFilter{
		Name:      req.Name,
		Key:       req.Key,
		Status:    req.Status,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		SortField: req.SortField,
		SortDesc:  req.SortDesc,
	}

	// 调用带筛选的查询方法
	models, total, err := c.RoleService.GetListWithFilter(ctx, req.Page, req.Size, filter)
	if err != nil {
		return
	}

	// 构建响应
	res = &v2.GetListRes{
		Total: total,
	}
	res.List = make([]*v2.RoleModel, 0, len(models))
	for _, model := range models {
		res.List = append(res.List, c.roleModelToV2(model))
	}
	return
}

func (c *ControllerV2) roleModelToV2(model *role.RoleModel) *v2.RoleModel {
	return &v2.RoleModel{
		Id:         model.RoleInfo.Id,
		Name:       model.RoleInfo.Name,
		Key:        model.RoleInfo.Key,
		ParentId:   model.TreeNode.Pid,
		ParentName: model.ParentName,
		Remark:     model.RoleInfo.Remark,
		Sort:       model.RoleInfo.Sort,
		Status:     model.RoleInfo.Status,
		CreatedAt:  model.RoleInfo.CreatedAt,
		UpdatedAt:  model.RoleInfo.UpdatedAt,
	}
}
