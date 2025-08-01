// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package user

import (
	"yclw/ygpay/api/user"
	"yclw/ygpay/internal/logic/member"
	"yclw/ygpay/internal/logic/menu"
	"yclw/ygpay/internal/logic/role"
)

type ControllerV1 struct {
	MemberService *member.Member
	MenuService   *menu.Menu
}

func NewV1() user.IUserV1 {
	return &ControllerV1{
		MemberService: member.MemberService,
		MenuService:   menu.MenuService,
	}
}

type ControllerV2 struct {
	MemberService *member.Member
	MenuService   *menu.Menu
	RoleService   *role.Role
}

func NewV2() user.IUserV2 {
	return &ControllerV2{
		MemberService: member.MemberService,
		MenuService:   menu.MenuService,
		RoleService:   role.RoleService,
	}
}
