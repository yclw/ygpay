package menu

import (
	"context"

	v2 "yclw/ygpay/api/menu/v2"
)

func (c *ControllerV2) Delete(ctx context.Context, req *v2.DeleteReq) (res *v2.DeleteRes, err error) {
	err = c.menuService.Delete(ctx, req.Id)
	if err != nil {
		return
	}
	return
}
