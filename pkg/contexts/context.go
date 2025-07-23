package contexts

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改
func Init(r *ghttp.Request, customCtx *Context) {
	r.SetCtxVar(ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func Get(ctx context.Context) *Context {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func SetUser(ctx context.Context, ctxUser *Identity) {
	Get(ctx).User = ctxUser
}

// SetResponse 设置组件响应 用于访问日志使用
func SetResponse(ctx context.Context, response *Response) {
	Get(ctx).Response = response
}

// GetUser 获取用户信息
func GetUser(ctx context.Context) *Identity {
	return Get(ctx).User
}

// GetUserId 获取用户ID
func GetUserId(ctx context.Context) int64 {
	return GetUser(ctx).Id
}

// GetRoleId 获取用户角色ID
func GetRoleId(ctx context.Context) int64 {
	return GetUser(ctx).RoleId
}
