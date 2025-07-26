package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/pkg/contexts"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	uid := contexts.GetUserUid(ctx)
	err = c.MemberService.Delete(ctx, uid)
	if err != nil {
		return
	}
	return
}
