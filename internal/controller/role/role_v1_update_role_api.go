package role

import (
	"context"
	"slices"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) UpdateRoleApi(ctx context.Context, req *v1.UpdateRoleApiReq) (res *v1.UpdateRoleApiRes, err error) {
	operator := contexts.GetRoleId(ctx)

	// 获取可用API，当前为操作角色API
	enabledApis, err := c.ApiService.GetRoleApi(ctx, operator)
	if err != nil {
		return
	}

	// 构建可用API map
	enabledApisMap := make(map[int64]bool)
	for _, api := range enabledApis {
		enabledApisMap[api.ApiInfo.Id] = true
	}

	// 过滤掉无权限的API
	req.ApiList = slices.DeleteFunc(req.ApiList, func(apiId int64) bool {
		return !enabledApisMap[apiId]
	})

	err = c.ApiService.UpdateRoleApi(ctx, req.Id, req.ApiList)
	return
}
