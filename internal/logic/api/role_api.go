package api

import (
	"context"
	"yclw/ygpay/internal/dao"
)

// GetSubList 获取子角色api列表
func (a *Api) GetSubList(ctx context.Context, roleId int64) (res []*ApiModel, err error) {

	// 获取子角色id列表
	subRoleIds, err := dao.RoleTree.FindSubRoleIds(ctx, roleId)
	if err != nil {
		return
	}

	// 获取子角色api列表
	apiIds, err := dao.RoleApi.FindRoleIdsByRoleIds(ctx, subRoleIds)
	if err != nil {
		return
	}

	// 获取api列表
	apis, err := dao.ApiInfo.FindByApiIds(ctx, apiIds)
	if err != nil {
		return
	}

	res = make([]*ApiModel, 0, len(apis))
	for _, api := range apis {
		res = append(res, &ApiModel{ApiInfo: api})
	}

	return
}
