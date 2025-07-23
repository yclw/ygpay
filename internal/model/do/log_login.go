// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LogLogin is the golang structure of table t_log_login for DAO operations like Where/Data.
type LogLogin struct {
	g.Meta        `orm:"table:t_log_login, do:true"`
	Id            interface{} // 登录日志ID
	MemberId      interface{} // 用户ID
	Username      interface{} // 登录账号
	IpAddress     interface{} // IP地址
	UserAgent     interface{} // 用户代理
	LoginLocation interface{} // 登录地点
	Browser       interface{} // 浏览器
	Os            interface{} // 操作系统
	LoginStatus   interface{} // 登录状态: 0失败 1成功
	LoginMessage  interface{} // 登录信息
	LoginTime     *gtime.Time // 登录时间
}
