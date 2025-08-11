package casbin

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/global"

	"github.com/gogf/gf/v2/errors/gerror"
)

var CasbinService = NewCasbin()

type Casbin struct {
}

func NewCasbin() *Casbin {
	return &Casbin{}
}

// SyncRoleApi 同步角色API权限到Casbin
func (c *Casbin) SyncRoleApi(ctx context.Context) (err error) {

	// 验证enforcer
	enforcer := global.Casbin()
	if enforcer == nil {
		return gerror.New("Casbin enforcer未初始化")
	}

	// 初始化策略
	rules := [][]string{}

	// 获取启用的角色列表
	roles, err := dao.RoleInfo.FindAllEnabled(ctx)
	if err != nil {
		return
	}

	// 获取角色关联的APIID列表
	for _, role := range roles {
		roleId := role.Id

		// 获取角色关联的APIID列表
		apiIds, err := dao.RoleApi.FindApiIdsByRoleId(ctx, roleId)
		if err != nil {
			return err
		}

		// 获取启用状态的API列表
		apis, err := dao.ApiInfo.FindEnabledByApiIds(ctx, apiIds)
		if err != nil {
			return err
		}

		// 添加策略
		for _, api := range apis {
			rules = append(rules, []string{role.Key, api.Path, api.Method})
		}
	}

	// 清空内存中的现有策略
	enforcer.ClearPolicy()

	// 添加策略到内存
	ok, err := enforcer.AddPolicies(rules)
	if err != nil || !ok {
		return
	}

	// 保存策略到数据库
	err = enforcer.SavePolicy()
	return
}

// RefreshEnforcer 刷新Casbin Enforcer
func (c *Casbin) RefreshEnforcer(ctx context.Context) (err error) {
	enforcer := global.Casbin()
	if enforcer == nil {
		return gerror.New("Casbin enforcer未初始化")
	}

	// 重新加载策略
	err = enforcer.LoadPolicy()
	return
}
