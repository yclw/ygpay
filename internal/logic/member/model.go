package member

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// MemberModel 用户模型
type MemberModel struct {
	*entity.MemberInfo
	*entity.MemberRole
	*MemberLoginStatModel
}

// MemberUpdateModel 用户更新模型
type MemberUpdateModel struct {
	Uid string `json:"uid"                 dc:"用户ID"`
	*do.MemberInfo
	*do.MemberRole
}

// MemberCreateModel 用户创建模型
type MemberCreateModel struct {
	Uid string `json:"uid"                 dc:"用户ID"`
	*do.MemberInfo
	*do.MemberRole
}

// MemberLoginStat 用户登录统计
type MemberLoginStatModel struct {
	LoginCount  int         `json:"loginCount"  dc:"登录次数"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
	LastLoginIp string      `json:"lastLoginIp" dc:"最后登录IP"`
}

// MemberListFilter 用户列表筛选参数
type MemberListFilter struct {
	Username  string      `json:"username"`  // 用户名
	Nickname  string      `json:"nickname"`  // 昵称
	Email     string      `json:"email"`     // 邮箱
	Mobile    string      `json:"mobile"`    // 手机号码
	RoleId    *int64      `json:"roleId"`    // 角色ID筛选
	Sex       *int        `json:"sex"`       // 性别筛选
	Status    *int        `json:"status"`    // 状态筛选
	StartDate *gtime.Time `json:"startDate"` // 开始日期
	EndDate   *gtime.Time `json:"endDate"`   // 结束日期
	SortField string      `json:"sortField"` // 排序字段
	SortDesc  bool        `json:"sortDesc"`  // 是否降序
}
