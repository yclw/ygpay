package member

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/global"
	"yclw/ygpay/internal/model/entity"
	"yclw/ygpay/pkg/token"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

var MemberService = NewMember()

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

// GetOne 获取单个用户信息
func (m *Member) GetOne(ctx context.Context, uid string) (res *MemberModel, err error) {
	res = &MemberModel{
		MemberRole: &entity.MemberRole{},
	}

	if uid == "" {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	member, err := dao.MemberInfo.FindByUid(ctx, uid)
	if err != nil {
		err = gerror.Wrap(err, "获取用户信息失败")
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
			MemberRole: &entity.MemberRole{},
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

// GetListWithFilter 获取用户列表（带分页和筛选）
func (s *Member) GetListWithFilter(ctx context.Context, page, size int, filter *MemberListFilter) (res []*MemberModel, total int, err error) {
	// 构建查询选项
	options := s.buildQueryOptions(filter)

	// 获取筛选后的用户信息
	members, total, err := dao.MemberInfo.FindWithPageAndOptions(ctx, page, size, options...)
	if err != nil {
		return
	}

	res = make([]*MemberModel, 0, len(members))
	for _, member := range members {
		memberModel := &MemberModel{
			MemberInfo: member,
			MemberRole: &entity.MemberRole{},
		}

		// 获取登录统计信息
		memberModel.MemberLoginStatModel, err = s.getLoginStat(ctx, member.Id)
		if err != nil {
			return nil, 0, err
		}

		// 获取角色ID
		roleId, err := dao.MemberRole.FindRoleIdByMemberId(ctx, member.Id)
		if err != nil {
			return nil, 0, err
		}
		memberModel.RoleId = roleId

		res = append(res, memberModel)
	}
	return
}

// buildQueryOptions 构建查询选项
func (s *Member) buildQueryOptions(filter *MemberListFilter) []dao.QueryOption {
	if filter == nil {
		return nil
	}

	var options []dao.QueryOption
	cols := dao.MemberInfo.Columns()

	// 用户名筛选
	if filter.Username != "" {
		options = append(options, dao.WhereLike(cols.Username, "%"+filter.Username+"%"))
	}

	// 昵称筛选
	if filter.Nickname != "" {
		options = append(options, dao.WhereLike(cols.Nickname, "%"+filter.Nickname+"%"))
	}

	// 邮箱筛选
	if filter.Email != "" {
		options = append(options, dao.WhereLike(cols.Email, "%"+filter.Email+"%"))
	}

	// 手机号筛选
	if filter.Mobile != "" {
		options = append(options, dao.WhereLike(cols.Mobile, "%"+filter.Mobile+"%"))
	}

	// 性别筛选
	if filter.Sex != nil {
		options = append(options, dao.Where(cols.Sex, *filter.Sex))
	}

	// 状态筛选
	if filter.Status != nil {
		options = append(options, dao.Where(cols.Status, *filter.Status))
	}

	// 日期范围筛选
	if filter.StartDate != nil || filter.EndDate != nil {
		if filter.StartDate != nil && filter.EndDate != nil {
			options = append(options, dao.WhereBetween(cols.CreatedAt, filter.StartDate, filter.EndDate))
		} else if filter.StartDate != nil {
			options = append(options, dao.Where(cols.CreatedAt+" >=", filter.StartDate))
		} else if filter.EndDate != nil {
			options = append(options, dao.Where(cols.CreatedAt+" <=", filter.EndDate))
		}
	}

	// 排序
	if filter.SortField != "" {
		if filter.SortDesc {
			options = append(options, dao.OrderDesc(filter.SortField))
		} else {
			options = append(options, dao.OrderAsc(filter.SortField))
		}
	} else {
		// 默认按创建时间降序
		options = append(options, dao.OrderDesc(cols.CreatedAt))
	}

	return options
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
