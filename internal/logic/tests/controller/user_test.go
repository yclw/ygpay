package controller

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	v1 "yclw/ygpay/api/user/v1"
	"yclw/ygpay/internal/cmd"
	"yclw/ygpay/internal/controller/user"
	"yclw/ygpay/pkg/contexts"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func TestGetRoleMenu(t *testing.T) {
	// 初始化应用配置和数据库连接
	ctx := context.Background()
	cmd.Init(ctx)

	// 创建自定义上下文
	customCtx := &contexts.Context{
		Session: nil,
		Data:    make(g.Map),
		User: &contexts.Identity{
			Uid:    "1",
			RoleId: 1,
		},
	}

	// 创建HTTP请求并设置上下文
	req := &ghttp.Request{Request: &http.Request{}}
	contexts.Init(req, customCtx)

	// 将自定义上下文设置到请求上下文中
	ctx = context.WithValue(ctx, contexts.ContextKey, customCtx)

	c := user.NewV1()

	input := &v1.GetUserMenuReq{}

	res, err := c.GetUserMenu(ctx, input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("res", res)
}
