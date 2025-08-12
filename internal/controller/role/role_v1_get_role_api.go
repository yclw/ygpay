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
	// 获取操作者角色UID
	operatorRoleUid := contexts.GetRoleUid(ctx)

	// 获得带Use标记的API列表
	enabledApis, err := c.ApiService.GetRoleApi(ctx, operatorRoleUid, req.RoleUid)
	if err != nil {
		return
	}

	// 转换为响应格式
	roleApis := c.convertToApiModels(enabledApis)
	c.sortApiModels(roleApis)

	res = &v1.GetRoleApiRes{
		ApiList: c.groupApis(roleApis),
	}
	return
}

// 转换API模型为响应格式
func (c *ControllerV1) convertToApiModels(apis []*api.ApiModel) []*v1.ApiModel {
	roleApis := make([]*v1.ApiModel, 0, len(apis))
	for _, api := range apis {
		roleApis = append(roleApis, &v1.ApiModel{
			ApiUid: api.ApiInfo.ApiUid,
			Path:   api.ApiInfo.Path,
			Method: api.ApiInfo.Method,
			Group:  api.ApiInfo.GroupName,
			Sort:   api.ApiInfo.Sort,
			Use:    api.Use,
		})
	}
	return roleApis
}

func (c *ControllerV1) sortApiModels(models []*v1.ApiModel) {
	slices.SortFunc(models, func(a, b *v1.ApiModel) int {
		return cmp.Compare(a.Sort, b.Sort)
	})
}

// API分组
func (c *ControllerV1) groupApis(apis []*v1.ApiModel) []*v1.ApiGroupModel {
	groupMap := make(map[string][]*v1.ApiModel)
	for _, api := range apis {
		groupMap[api.Group] = append(groupMap[api.Group], api)
	}

	groupList := make([]*v1.ApiGroupModel, 0, len(groupMap))
	for groupName, apiList := range groupMap {
		groupList = append(groupList, &v1.ApiGroupModel{
			GroupName: groupName,
			Children:  apiList,
		})
	}
	return groupList
}
