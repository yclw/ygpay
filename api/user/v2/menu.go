package v2

import "github.com/gogf/gf/v2/frame/g"

// 参考 https://pure-admin.cn/pages/routerMenu

// 菜单
type UserMenu struct {
	Type      int          `json:"type" dc:"类型（0:目录 1:菜单 2:按钮）"`
	Name      string       `json:"name" dc:"名称"`
	Path      string       `json:"path" dc:"路径"`
	Component string       `json:"component,omitempty" dc:"组件路径（内链返回）"`
	Redirect  string       `json:"redirect,omitempty" dc:"重定向（目录返回）"`
	FrameSrc  string       `json:"frameSrc,omitempty" dc:"内嵌地址（外链时有效）"`
	Children  []UserMenu   `json:"children,omitempty" dc:"子菜单"`
	Meta      UserMenuMeta `json:"meta" dc:"元数据"`
}

// 菜单元数据
type UserMenuMeta struct {
	Icon       string `json:"icon,omitempty" dc:"图标"`
	Title      string `json:"title" dc:"标题"`
	Rank       int64  `json:"rank" dc:"排序"`
	ShowParent bool   `json:"showParent" dc:"是否显示父菜单"`
	KeepAlive  bool   `json:"keepAlive" dc:"是否缓存"`
	ShowLink   bool   `json:"showLink" dc:"是否显示链接"`
}

// 获取用户菜单树
type GetUserMenuReq struct {
	g.Meta `path:"/user/menu" method:"get" tags:"用户" summary:"获取用户菜单树"`
}

type GetUserMenuRes struct {
	Menu []UserMenu `json:"menu" dc:"菜单树"`
}
