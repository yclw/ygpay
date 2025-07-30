package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/api"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetRoleApi(ctx context.Context, req *v1.GetRoleApiReq) (res *v1.GetRoleApiRes, err error) {

	// 获取角色可用api列表
	apiList, err := c.ApiService.GetRoleApiList(ctx, contexts.GetRoleId(ctx))
	if err != nil {
		return
	}

	// 获取使用角色api列表
	useApiList, err := c.ApiService.GetRoleApiList(ctx, req.Id)
	if err != nil {
		return
	}

	// 获取当前用户api列表
	res = &v1.GetRoleApiRes{
		ApiList: c.apiModelToV1List(ctx, apiList, useApiList),
	}

	return
}

func (c *ControllerV1) apiModelToV1List(ctx context.Context, apiList []*api.ApiModel, useApiList []*api.ApiModel) []v1.RoleApiModel {

	// 获取使用角色api列表
	useApiMap := make(map[int64]bool)
	for _, api := range useApiList {
		useApiMap[api.ApiInfo.Id] = true
	}

	res := make([]v1.RoleApiModel, 0, len(apiList))
	for _, api := range apiList {
		_, exist := useApiMap[api.ApiInfo.Id]
		res = append(res, v1.RoleApiModel{
			Id:        api.ApiInfo.Id,
			Name:      api.ApiInfo.Name,
			Path:      api.ApiInfo.Path,
			Method:    api.ApiInfo.Method,
			GroupName: api.ApiInfo.GroupName,
			Use:       exist,
		})
	}
	return res
}
