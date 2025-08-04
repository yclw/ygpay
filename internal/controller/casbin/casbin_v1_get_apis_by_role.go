package casbin

import (
	"context"

	v1 "yclw/ygpay/api/casbin/v1"
)

func (c *ControllerV1) GetApisByRole(ctx context.Context, req *v1.GetApisByRoleReq) (res *v1.GetApisByRoleRes, err error) {
	casbinApis, err := c.CasbinService.GetApisByRole(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}

	// 转换为API模型
	apiModels := make([]v1.ApiModel, 0, len(casbinApis))
	for _, casbinApi := range casbinApis {
		apiModels = append(apiModels, v1.ApiModel{
			Path:   casbinApi.Path,
			Method: casbinApi.Method,
		})
	}

	return &v1.GetApisByRoleRes{
		Apis: apiModels,
	}, nil
}
