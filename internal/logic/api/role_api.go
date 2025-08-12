package api

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/util/slices"
)

// GetRoleApi 获取角色API分配数据（带Use标记的API列表）
func (a *Api) GetRoleApi(ctx context.Context, operatorRoleUid, targetRoleUid string) (res []*ApiModel, err error) {

	// 获取操作者角色ID可用的API列表
	operatorRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, operatorRoleUid)
	if err != nil {
		return nil, err
	}
	enabledApiIds, err := dao.RoleApi.FindApiIdsByRoleId(ctx, operatorRoleId)
	if err != nil {
		return nil, err
	}
	enabledApis, err := dao.ApiInfo.FindEnabledByApiIds(ctx, enabledApiIds)
	if err != nil {
		return nil, err
	}

	// 获取目标角色已用的API ID列表
	usedApiIds := enabledApiIds
	if operatorRoleUid != targetRoleUid {
		// 获取目标角色已用的API ID列表
		targetRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, targetRoleUid)
		if err != nil {
			return nil, err
		}
		usedApiIds, err = dao.RoleApi.FindApiIdsByRoleId(ctx, targetRoleId)
		if err != nil {
			return nil, err
		}
	}

	usedApiMap := make(map[int64]bool, len(usedApiIds))
	for _, id := range usedApiIds {
		usedApiMap[id] = true
	}

	// 转换为ApiModel
	res = make([]*ApiModel, 0, len(enabledApis))
	for _, api := range enabledApis {
		res = append(res, &ApiModel{
			ApiInfo: api,
			Use:     usedApiMap[api.Id],
		})
	}

	return
}

// UpdateRoleApi 更新角色AP
func (a *Api) UpdateRoleApi(ctx context.Context, operatorRoleUid, targetRoleUid string, apiUids []string) (err error) {

	// 获取目标角色ID
	targetRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, targetRoleUid)
	if err != nil {
		return
	}

	// 将API UID列表转换为ID列表
	apiIds, err := dao.ApiInfo.FindIdsByApiUids(ctx, apiUids)
	if err != nil {
		return
	}

	// 权限过滤：只允许操作者有权限的API
	operatorRoleId, err := dao.RoleInfo.FindIdByRoleUid(ctx, operatorRoleUid)
	if err != nil {
		return
	}
	filteredApiIds, err := a.FilterAllowedApiIdsByRole(ctx, apiIds, operatorRoleId)
	if err != nil {
		return
	}

	// 更新角色API
	_, err = dao.RoleApi.UpdateRoleApi(ctx, targetRoleId, filteredApiIds)
	return
}

// FilterAllowedApiIdsByRole 根据角色过滤允许的API列表，直接返回ID
func (a *Api) FilterAllowedApiIdsByRole(ctx context.Context, apiIds []int64, roleId int64) ([]int64, error) {

	// 获取角色允许的API ID列表
	allowedApiIds, err := dao.RoleApi.FindApiIdsByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}

	// 计算交集（只保留既在请求列表中又被角色允许的API）
	return slices.IntersectInt64s(apiIds, allowedApiIds), nil
}
