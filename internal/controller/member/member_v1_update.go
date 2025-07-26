package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/internal/logic/member"
	"yclw/ygpay/util/encrypt"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	input := &member.MemberUpdateModel{}
	input.Uid = req.Uid
	input.Username = req.Username
	input.Avatar = req.Avatar
	input.Sex = req.Sex
	input.Email = req.Email
	input.Mobile = req.Mobile
	input.Address = req.Address
	input.Remark = req.Remark
	input.Sort = req.Sort
	input.Status = req.Status
	input.RoleId = req.RoleId
	input.PasswordHash, err = encrypt.HashPassword(req.Password, encrypt.DefaultCost)
	if err != nil {
		err = gerror.Wrap(err, "密码加密失败")
		return
	}

	err = c.MemberService.Update(ctx, input)
	if err != nil {
		return
	}
	return
}
