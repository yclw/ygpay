package v2

import (
	"yclw/ygpay/util/tree"

	"github.com/gogf/gf/v2/frame/g"
)

type MenuTree = tree.TreeNode

// GetListReq 获取菜单列表
type GetTreeReq struct {
	g.Meta `path:"/menu/tree" method:"get" tags:"菜单管理" summary:"获取菜单列表"`
}

type GetTreeRes struct {
	MenuTree []*MenuTree `json:"menuTree" dc:"菜单树"`
}
