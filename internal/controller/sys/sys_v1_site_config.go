package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"yclw/ygpay/api/sys/v1"
)

func (c *ControllerV1) SiteConfig(ctx context.Context, req *v1.SiteConfigReq) (res *v1.SiteConfigRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
