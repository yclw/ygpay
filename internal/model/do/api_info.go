// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApiInfo is the golang structure of table t_api_info for DAO operations like Where/Data.
type ApiInfo struct {
	g.Meta      `orm:"table:t_api_info, do:true"`
	Id          interface{} // API ID
	Name        interface{} // API名称
	Path        interface{} // API路径
	Method      interface{} // API方法
	GroupName   interface{} // API分组
	Description interface{} // API描述
	NeedAuth    interface{} // 是否需要认证: 0否 1是
	RateLimit   interface{} // 限流次数/分钟
	Sort        interface{} // 排序
	Status      interface{} // 状态: 0禁用 1启用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
