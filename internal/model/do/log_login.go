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
	g.Meta     `orm:"table:t_log_login, do:true"`
	Id         interface{} // 日志ID
	ReqId      interface{} // 请求ID
	MemberId   interface{} // 用户ID
	Username   interface{} // 用户名
	Response   interface{} // 响应数据
	LoginAt    *gtime.Time // 登录时间
	LoginIp    interface{} // 登录IP
	ProvinceId interface{} // 省编码
	CityId     interface{} // 市编码
	UserAgent  interface{} // UA信息
	ErrMsg     interface{} // 错误提示
	Status     interface{} // 状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
