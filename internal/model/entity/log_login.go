// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LogLogin is the golang structure for table log_login.
type LogLogin struct {
	Id            int64       `json:"id"            orm:"id"             ` // 登录日志ID
	MemberId      int64       `json:"memberId"      orm:"member_id"      ` // 用户ID
	Username      string      `json:"username"      orm:"username"       ` // 登录账号
	IpAddress     string      `json:"ipAddress"     orm:"ip_address"     ` // IP地址
	UserAgent     string      `json:"userAgent"     orm:"user_agent"     ` // 用户代理
	LoginLocation string      `json:"loginLocation" orm:"login_location" ` // 登录地点
	Browser       string      `json:"browser"       orm:"browser"        ` // 浏览器
	Os            string      `json:"os"            orm:"os"             ` // 操作系统
	LoginStatus   int         `json:"loginStatus"   orm:"login_status"   ` // 登录状态: 0失败 1成功
	LoginMessage  string      `json:"loginMessage"  orm:"login_message"  ` // 登录信息
	LoginTime     *gtime.Time `json:"loginTime"     orm:"login_time"     ` // 登录时间
}
