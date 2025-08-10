// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table t_sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta      `orm:"table:t_sys_config, do:true"`
	Id          interface{} // 配置ID
	Group       interface{} // 配置分组
	Key         interface{} // 参数键名
	Value       interface{} // 参数值
	Description interface{} //
	Sort        interface{} // 排序
	Status      interface{} // 状态: 0禁用 1启用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
