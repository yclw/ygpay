package menu

import (
	"context"

	v1 "yclw/ygpay/api/menu/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = c.menuService.Delete(ctx, req.Id)
	if err != nil {
		return
	}
	return
}
