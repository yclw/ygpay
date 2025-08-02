package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = c.MemberService.Delete(ctx, req.Uid)
	if err != nil {
		return
	}
	return
}
