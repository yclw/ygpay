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
