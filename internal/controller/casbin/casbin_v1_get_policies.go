package casbin

import (
	"context"

	v1 "yclw/ygpay/api/casbin/v1"
)

func (c *ControllerV1) GetPolicies(ctx context.Context, req *v1.GetPoliciesReq) (res *v1.GetPoliciesRes, err error) {
	casbinPolicies, err := c.CasbinService.GetPolicies(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}

	// 转换为API模型
	policyModels := make([]v1.PolicyModel, 0, len(casbinPolicies))
	for _, casbinPolicy := range casbinPolicies {
		policyModels = append(policyModels, v1.PolicyModel{
			RoleId: casbinPolicy.RoleId,
			Path:   casbinPolicy.Path,
			Method: casbinPolicy.Method,
		})
	}

	return &v1.GetPoliciesRes{
		Policies: policyModels,
	}, nil
}
