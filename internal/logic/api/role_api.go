package api

import (
	"context"
	"yclw/ygpay/internal/dao"
)

// GetRoleApi 获取角色api列表
func (a *Api) GetRoleApi(ctx context.Context, roleId int64) (res []*ApiModel, err error) {
	// 获取角色apiID列表
	apiIds, err := dao.RoleApi.FindApiIdsByRoleId(ctx, roleId)
	if err != nil {
		return
	}

	// 获取角色api信息
	apis, err := dao.ApiInfo.FindEnabledByApiIds(ctx, apiIds)
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

// UpdateRoleApi 更新角色API
func (a *Api) UpdateRoleApi(ctx context.Context, roleId int64, apiUids []string) (err error) {
	// 获取apiId列表
	apiIds, err := dao.ApiInfo.FindIdsByApiUids(ctx, apiUids)
	if err != nil {
		return
	}

	//TODO: 考虑事务

	// 删除角色API
	_, err = dao.RoleApi.DeleteByRoleId(ctx, roleId)
	if err != nil {
		return
	}

	// 添加角色API
	_, err = dao.RoleApi.AddRoleApi(ctx, roleId, apiIds)
	return
}

// AddRoleApi 添加角色API
func (a *Api) AddRoleApi(ctx context.Context, roleId int64, apiUid string) (err error) {
	// 获取apiId
	apiId, err := dao.ApiInfo.FindIdByApiUid(ctx, apiUid)
	if err != nil {
		return
	}

	// 添加角色API
	_, err = dao.RoleApi.AddRoleApi(ctx, roleId, []int64{apiId})
	return
}
