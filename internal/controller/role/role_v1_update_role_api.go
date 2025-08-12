package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) UpdateRoleApi(ctx context.Context, req *v1.UpdateRoleApiReq) (res *v1.UpdateRoleApiRes, err error) {
	// 获取操作者角色UID
	operatorRoleUid := contexts.GetRoleUid(ctx)

	err = c.ApiService.UpdateRoleApi(ctx, operatorRoleUid, req.RoleUid, req.ApiList)
	return
}
