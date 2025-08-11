package role

import (
	"context"
	"slices"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) UpdateRoleApi(ctx context.Context, req *v1.UpdateRoleApiReq) (res *v1.UpdateRoleApiRes, err error) {
	operatorRoleUid := contexts.GetRoleUid(ctx)

	// 根据操作者RoleUid获取RoleId
	operatorRoleId, err := c.RoleService.GetRoleIdByUid(ctx, operatorRoleUid)
	if err != nil {
		return
	}

	// 获取可用API，当前为操作角色API
	enabledApis, err := c.ApiService.GetRoleApi(ctx, operatorRoleId)
	if err != nil {
		return
	}

	// 构建可用API map
	enabledApisMap := make(map[string]bool)
	for _, api := range enabledApis {
		enabledApisMap[api.ApiInfo.ApiUid] = true
	}

	// 过滤掉无权限的API
	req.ApiList = slices.DeleteFunc(req.ApiList, func(apiUid string) bool {
		return !enabledApisMap[apiUid]
	})

	// 根据目标roleUid获取roleId
	targetRoleId, err := c.RoleService.GetRoleIdByUid(ctx, req.RoleUid)
	if err != nil {
		return
	}

	err = c.ApiService.UpdateRoleApi(ctx, targetRoleId, req.ApiList)
	return
}
