package v1

import "github.com/gogf/gf/v2/frame/g"

// 参考 https://pure-admin.cn/pages/routerMenu

// 菜单
type UserMenu struct {
	Name              string       `json:"name" dc:"名称"`
	Path              string       `json:"path" dc:"路径"`
	Component         string       `json:"component" dc:"组件路径"`
	NoShowingChildren bool         `json:"noShowingChildren" dc:"是否显示子菜单"`
	Children          []UserMenu   `json:"children" dc:"子菜单"`
	Value             string       `json:"value" dc:"值"`
	Meta              UserMenuMeta `json:"meta" dc:"元数据"`
	ShowTooltip       bool         `json:"showTooltip" dc:"是否显示提示"`
	ParentId          int64        `json:"parentId" dc:"父ID"`
	// PathList          []int64      `json:"pathList" dc:"路径列表"`
	Redirect string `json:"redirect" dc:"重定向"`
}

// 菜单元数据
type UserMenuMeta struct {
	Icon       string `json:"icon" dc:"图标"`
	Title      string `json:"title" dc:"标题"`
	Rank       int64  `json:"rank" dc:"排序"`
	ShowParent bool   `json:"showParent" dc:"是否显示父菜单"`

	Auths []string `json:"auths" dc:"控件权限"`
}

// 获取用户菜单树
type GetUserMenuReq struct {
	g.Meta `path:"/user/menu" method:"get" tags:"用户" summary:"获取用户菜单树"`
}

type GetUserMenuRes struct {
	Menu []UserMenu `json:"menu" dc:"菜单树"`
}
