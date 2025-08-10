package role

import (
	"cmp"
	"context"
	"slices"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/internal/logic/api"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) GetRoleApi(ctx context.Context, req *v1.GetRoleApiReq) (res *v1.GetRoleApiRes, err error) {
	operator := contexts.GetRoleId(ctx)

	// 获取可用API，当前为操作角色API
	enabledApis, err := c.ApiService.GetRoleApi(ctx, operator)
	if err != nil {
		return
	}

	// 获取已用API
	usedApis, err := c.ApiService.GetRoleApi(ctx, req.Id)
	if err != nil {
		return
	}

	// 构建已用API map
	usedApisMap := make(map[int64]bool)
	for _, api := range usedApis {
		usedApisMap[api.ApiInfo.Id] = true
	}

	// 转换可用API为v1.ApiModel
	roleApis := make([]*v1.ApiModel, 0)
	for _, api := range enabledApis {
		roleApi := c.apiModelToRoleApi(api, usedApisMap)
		roleApis = append(roleApis, roleApi)
	}

	// 排序
	slices.SortFunc(roleApis, func(a, b *v1.ApiModel) int {
		return cmp.Compare(a.Sort, b.Sort)
	})

	// 按照Group分组
	groupList := c.groupApis(roleApis)

	res = &v1.GetRoleApiRes{
		ApiList: groupList,
	}

	return
}

// apiModelToRoleApi 将apiModel转换为roleApiModel
func (c *ControllerV1) apiModelToRoleApi(apiModel *api.ApiModel, usedApisMap map[int64]bool) *v1.ApiModel {
	roleApi := v1.ApiModel{
		ApiUid: apiModel.ApiInfo.ApiUid,
		Path:   apiModel.ApiInfo.Path,
		Method: apiModel.ApiInfo.Method,
		Group:  apiModel.ApiInfo.GroupName,
		Sort:   apiModel.ApiInfo.Sort,
		Use:    usedApisMap[apiModel.ApiInfo.Id],
	}
	return &roleApi
}

// groupApis 按照Group分组api
func (c *ControllerV1) groupApis(apis []*v1.ApiModel) []*v1.ApiGroupModel {
	groupMap := make(map[string][]*v1.ApiModel)
	for _, api := range apis {
		groupMap[api.Group] = append(groupMap[api.Group], api)
	}

	groupList := make([]*v1.ApiGroupModel, 0)
	for groupName, apiList := range groupMap {
		groupList = append(groupList, &v1.ApiGroupModel{
			GroupName: groupName,
			Children:  apiList,
		})
	}
	return groupList
}
