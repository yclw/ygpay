package role

import (
	"context"

	v1 "yclw/ygpay/api/role/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error) {
	// 获取操作者角色UID
	operatorRoleUid := contexts.GetRoleUid(ctx)

	err = c.MenuService.UpdateRoleMenu(ctx, operatorRoleUid, req.RoleUid, req.MenuList)
	return
}
