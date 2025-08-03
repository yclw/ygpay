package member

import (
	"context"

	"github.com/gogf/gf/v2/text/gstr"
)

// GetOneEncrypt 获取单个用户信息--加密处理
func (m *Member) GetOneEncrypt(ctx context.Context, uid string) (res *MemberModel, err error) {
	res, err = m.GetOne(ctx, uid)
	if err != nil {
		return
	}
	res.Mobile = gstr.HideStr(res.Mobile, 40, `*`) // 手机号脱敏
	res.Email = gstr.HideStr(res.Email, 40, `*`)   // 邮箱脱敏
	return
}
