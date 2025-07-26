package member

import (
	"context"
	"yclw/ygpay/internal/dao"

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

// GetSubList 获取子用户列表
func (s *Member) GetSubList(ctx context.Context, roleId int64) (res []*MemberModel, err error) {

	// 获取子角色id列表
	subRoleIds, err := dao.RoleTree.FindSubRoleIds(ctx, roleId)
	if err != nil {
		return
	}
	// 获取用户ID列表
	userIds, err := dao.MemberRole.FindUserIdsByRoleIds(ctx, subRoleIds)
	if err != nil {
		return
	}

	// 获取子用户列表
	members, err := dao.MemberInfo.FindByUserIds(ctx, userIds)
	if err != nil {
		return
	}

	res = make([]*MemberModel, 0, len(members))
	for _, member := range members {
		memberModel := &MemberModel{
			MemberInfo: member,
		}

		// 获取登录统计信息
		memberModel.MemberLoginStatModel, err = s.getLoginStat(ctx, member.Id)
		if err != nil {
			return nil, err
		}

		// 获取角色ID
		roleId, err := dao.MemberRole.FindRoleIdByMemberId(ctx, member.Id)
		if err != nil {
			return nil, err
		}
		memberModel.RoleId = roleId

		res = append(res, memberModel)
	}
	return
}
