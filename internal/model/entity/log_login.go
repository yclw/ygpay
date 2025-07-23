// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LogLogin is the golang structure for table log_login.
type LogLogin struct {
	Id         int64       `json:"id"         orm:"id"          ` // 日志ID
	ReqId      string      `json:"reqId"      orm:"req_id"      ` // 请求ID
	MemberId   int64       `json:"memberId"   orm:"member_id"   ` // 用户ID
	Username   string      `json:"username"   orm:"username"    ` // 用户名
	Response   string      `json:"response"   orm:"response"    ` // 响应数据
	LoginAt    *gtime.Time `json:"loginAt"    orm:"login_at"    ` // 登录时间
	LoginIp    string      `json:"loginIp"    orm:"login_ip"    ` // 登录IP
	ProvinceId int64       `json:"provinceId" orm:"province_id" ` // 省编码
	CityId     int64       `json:"cityId"     orm:"city_id"     ` // 市编码
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  ` // UA信息
	ErrMsg     string      `json:"errMsg"     orm:"err_msg"     ` // 错误提示
	Status     int         `json:"status"     orm:"status"      ` // 状态
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` // 修改时间
}
