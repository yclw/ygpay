package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/internal/consts"
	"yclw/ygpay/internal/logic/member"
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/util/encrypt"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/guid"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	model, err := c.createReqToCreateModel(req)
	if err != nil {
		return
	}

	err = c.MemberService.Create(ctx, model)
	if err != nil {
		return
	}

	res = &v1.CreateRes{}
	return
}

// createReqToCreateModel 将创建请求转换为创建模型
func (c *ControllerV1) createReqToCreateModel(req *v1.CreateReq) (*member.MemberCreateModel, error) {
	// 解码密码
	password, err := gbase64.Decode([]byte(req.Password))
	if err != nil {
		return nil, gerror.Wrap(err, "密码解析失败")
	}
	// 解密密码
	password, err = encrypt.AesECBDecrypt(password, []byte(consts.RequestEncryptKey))
	if err != nil {
		return nil, gerror.Wrap(err, "密码解析失败")
	}
	// 密码加密
	passwordHash, err := encrypt.HashPassword(string(password), encrypt.DefaultCost)
	if err != nil {
		return nil, gerror.Wrap(err, "密码加密失败")
	}

	return &member.MemberCreateModel{
		Uid: guid.S(),
		MemberInfo: &do.MemberInfo{
			Username:     req.Username,
			Nickname:     req.Nickname,
			PasswordHash: passwordHash,
			Avatar:       req.Avatar,
			Sex:          req.Sex,
			Email:        req.Email,
			Mobile:       req.Mobile,
			Address:      req.Address,
			Remark:       req.Remark,
			Sort:         req.Sort,
			Status:       req.Status,
		},
		MemberRole: &do.MemberRole{
			RoleId: req.RoleId,
		},
	}, nil
}
