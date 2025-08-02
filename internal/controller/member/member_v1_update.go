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
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	model, err := c.updateReqToUpdateModel(req)
	if err != nil {
		return
	}

	err = c.MemberService.Update(ctx, model)
	if err != nil {
		return
	}

	res = &v1.UpdateRes{}
	return
}

// updateReqToUpdateModel 将更新请求转换为更新模型
func (c *ControllerV1) updateReqToUpdateModel(req *v1.UpdateReq) (*member.MemberUpdateModel, error) {
	memberInfo := &do.MemberInfo{
		Uid:      req.Uid,
		Username: req.Username,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Sex:      req.Sex,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Address:  req.Address,
		Remark:   req.Remark,
		Sort:     req.Sort,
		Status:   req.Status,
	}

	// 密码安全处理：只有提供了新密码才更新，否则保留原密码
	if req.Password != "" {
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
		passwordHash, err := encrypt.HashPassword(string(password), encrypt.DefaultCost)
		if err != nil {
			return nil, gerror.Wrap(err, "密码加密失败")
		}
		memberInfo.PasswordHash = passwordHash
	}
	// 注意：如果Password为空，则不设置PasswordHash字段，
	// 配合DAO层的OmitNil()，确保数据库中的密码不会被清空

	return &member.MemberUpdateModel{
		Uid:        req.Uid,
		MemberInfo: memberInfo,
		MemberRole: &do.MemberRole{
			RoleId: req.RoleId,
		},
	}, nil
}
