package menu

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
)

type UserMenuModel struct {
	*entity.MenuInfo
	*entity.MenuTree
}

type MenuModel struct {
	*entity.MenuInfo
	*entity.MenuTree
}

type MenuCreateModel struct {
	Id int64
	*do.MenuInfo
	*do.MenuTree
}

type MenuUpdateModel struct {
	Id int64
	*do.MenuInfo
	*do.MenuTree
}
