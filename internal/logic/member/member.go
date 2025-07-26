package member

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/global"
	"yclw/ygpay/pkg/token"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

var MemberService = NewMember()

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

// GetOne 获取单个用户信息
func (m *Member) GetOne(ctx context.Context, uid string) (res *MemberModel, err error) {
	res = &MemberModel{}

	if uid == "" {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	member, err := dao.MemberInfo.FindByUid(ctx, uid)
	if err != nil {
		err = gerror.Wrap(err, "获取用户信息失败")
		return
	}
	if member == nil {
		err = gerror.New("用户不存在！")
		return
	}
	res.MemberInfo = member

	// 获取登录信息
	stat, err := m.getLoginStat(ctx, member.Id)
	if err != nil || stat == nil {
		return
	}
	res.MemberLoginStatModel = stat

	// 获取角色ID
	roleId, err := dao.MemberRole.FindRoleIdByMemberId(ctx, member.Id)
	if err != nil {
		return
	}
	res.RoleId = roleId

	return
}

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

// getLoginStat 获取用户登录统计
func (s *Member) getLoginStat(ctx context.Context, memberId int64) (res *MemberLoginStatModel, err error) {
	res = &MemberLoginStatModel{
		LastLoginAt: gtime.Now(),
	}
	last, err := dao.LogLogin.FindLastByMemberId(ctx, memberId)
	if err != nil {
		return
	}
	if last == nil {
		return
	}
	if last.LoginTime != nil {
		res.LastLoginAt = last.LoginTime
	}
	res.LastLoginIp = last.IpAddress
	res.LoginCount, err = dao.LogLogin.GetLoginCount(ctx, memberId)
	return
}

// GetAllList 获取所有用户列表
func (s *Member) GetAllList(ctx context.Context) (res []*MemberModel, err error) {
	members, err := dao.MemberInfo.FindAll(ctx)
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

// GetSubList 获取子用户列表
func (s *Member) GetSubList(ctx context.Context, memberId int64) (res []*MemberModel, err error) {
	// 获取用户角色
	roleId, err := dao.MemberRole.FindRoleIdByMemberId(ctx, memberId)
	if err != nil {
		return
	}
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

// Create 创建用户
func (s *Member) Create(ctx context.Context, req *MemberCreateModel) (err error) {
	// 创建用户
	id, err := dao.MemberInfo.Create(ctx, req.MemberInfo)
	if err != nil {
		err = gerror.Wrap(err, "创建用户失败")
		return
	}

	// 创建用户角色关系
	req.MemberRole.MemberId = id
	err = dao.MemberRole.Create(ctx, req.MemberRole)
	if err != nil {
		err = gerror.Wrap(err, "创建用户角色关系失败")
		return
	}
	return
}

// Update 更新用户
func (s *Member) Update(ctx context.Context, req *MemberUpdateModel) (err error) {
	if err = dao.MemberInfo.Update(ctx, req.MemberInfo); err != nil {
		err = gerror.Wrap(err, "更新用户失败")
		return
	}

	// 更新用户角色关系
	memberId, err := dao.MemberInfo.FindIdByUid(ctx, req.Uid)
	if err != nil {
		err = gerror.Wrap(err, "更新用户角色失败")
		return
	}
	req.MemberRole.MemberId = memberId
	err = dao.MemberRole.UpdateRoleIdByMemberId(ctx, req.MemberRole)
	if err != nil {
		err = gerror.Wrap(err, "更新用户角色失败")
		return
	}

	// 删除刷新token缓存
	refreshKey := token.RefreshJwt.GetCacheKey(ctx, req.Uid)
	global.Cache().Remove(ctx, refreshKey)
	return
}

// Delete 删除用户
func (s *Member) Delete(ctx context.Context, uid string) (err error) {
	if err = dao.MemberInfo.DeleteByUid(ctx, uid); err != nil {
		err = gerror.Wrap(err, "删除用户失败")
		return
	}
	// 删除刷新token缓存
	refreshKey := token.RefreshJwt.GetCacheKey(ctx, uid)
	global.Cache().Remove(ctx, refreshKey)
	return
}
