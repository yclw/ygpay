package contexts

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

const (
	ContextKey = "ContextKey" // 上下文变量存储键名，前后端系统共享
)

// Context 请求上下文结构
type Context struct {
	Session  *ghttp.Session // 会话
	User     *Identity      // 上下文用户信息
	Response *Response      // 请求响应
	Data     g.Map          // 自定kv变量
}

// Identity 上下文用户身份模型
type Identity struct {
	Id      int64       `json:"id"              dc:"用户ID"`
	RoleId  int64       `json:"roleId"          dc:"角色ID"`
	LoginAt *gtime.Time `json:"loginAt"         dc:"登录时间"`
}

// Response HTTP响应
type Response struct {
	Code      int         `json:"code"               dc:"状态码"`
	Message   string      `json:"message,omitempty"  dc:"提示消息"`
	Data      interface{} `json:"data,omitempty"     dc:"数据集"`
	Error     interface{} `json:"error,omitempty"    dc:"错误信息"`
	Timestamp int64       `json:"timestamp"          dc:"服务器时间戳"`
	TraceID   string      `json:"traceID" v:"0"      dc:"链路ID"`
}
