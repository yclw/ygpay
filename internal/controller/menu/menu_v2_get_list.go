package menu

import (
	"context"

	v2 "yclw/ygpay/api/menu/v2"
	"yclw/ygpay/internal/logic/menu"
)

func (c *ControllerV2) GetList(ctx context.Context, req *v2.GetListReq) (res *v2.GetListRes, err error) {
	// 构建筛选参数
	filter := &menu.MenuListFilter{
		Status: req.Status,
	}

	// 调用带筛选的查询方法
	models, total, err := c.menuService.GetListWithFilter(ctx, req.Page, req.Size, filter)
	if err != nil {
		return
	}

	// 构建响应
	res = &v2.GetListRes{
		Total: total,
	}
	res.List = make([]*v2.MenuModel, 0, len(models))
	for _, model := range models {
		res.List = append(res.List, c.menuModelToV2(model))
	}
	return
}
