package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/internal/logic/member"
	"yclw/ygpay/util/encrypt"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	input := &member.MemberCreateModel{}
	input.RoleId = req.RoleId
	input.Username = req.Username
	input.Nickname = req.Nickname
	input.RoleId = req.RoleId
	input.Avatar = req.Avatar
	input.Sex = req.Sex
	input.Email = req.Email
	input.Mobile = req.Mobile
	input.Address = req.Address
	input.Remark = req.Remark
	input.Sort = req.Sort
	input.Status = req.Status

	input.PasswordHash, err = encrypt.HashPassword(req.Password, encrypt.DefaultCost)
	if err != nil {
		err = gerror.Wrap(err, "密码加密失败")
		return
	}

	err = c.MemberService.Create(ctx, input)
	if err != nil {
		return
	}
	return
}
