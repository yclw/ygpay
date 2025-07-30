package api

import (
	"context"
	"yclw/ygpay/internal/dao"
)

// GetRoleApiList 获取角色api列表
func (a *Api) GetRoleApiList(ctx context.Context, roleId int64) (res []*ApiModel, err error) {

	// 获取角色api列表
	apiIds, err := dao.RoleApi.FindApiIdsByRoleId(ctx, roleId)
	if err != nil {
		return
	}

	// 获取api列表
	apis, err := dao.ApiInfo.FindByApiIds(ctx, apiIds)
	if err != nil {
		return
	}

	// 转换为ApiModel
	res = make([]*ApiModel, 0, len(apis))
	for _, api := range apis {
		res = append(res, &ApiModel{ApiInfo: api})
	}

	return
}

// SyncCasbin 同步Casbin
func (a *Api) SyncCasbin(ctx context.Context) (err error) {
	return
}
