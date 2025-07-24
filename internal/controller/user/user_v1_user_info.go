package user

import (
	"context"
	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/pkg/contexts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) UserInfo(ctx context.Context, _ *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	uid := contexts.GetUserUid(ctx)
	member, err := c.MemberService.GetOneEncrypt(ctx, uid)
	if err != nil {
		return
	}
	err = gconv.Struct(member, &res)
	if err != nil {
		err = gerror.Wrap(err, "字段复制失败")
		return
	}
	return
}
