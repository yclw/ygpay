package casbin

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"

	"yclw/ygpay/internal/consts"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/global"
	"yclw/ygpay/internal/model/entity"
)

type Casbin struct {
}

var CasbinService = &Casbin{}

// SyncRoleApi 同步角色API到Casbin表
func (c *Casbin) SyncRoleApi(ctx context.Context) error {
	enforcer := global.Casbin()
	if enforcer == nil {
		return gerror.New("Casbin enforcer未初始化")
	}

	// 清空所有casbin策略
	enforcer.ClearPolicy()

	// 获取所有角色信息（只获取启用状态的角色）
	cols := dao.RoleInfo.Columns()
	roles, err := dao.RoleInfo.Ctx(ctx).Where(cols.Status, consts.StatusEnabled).All()
	if err != nil {
		return gerror.Wrap(err, "获取角色信息失败")
	}

	// 遍历每个角色，同步其API权限
	for _, role := range roles {
		roleEntity := &entity.RoleInfo{}
		if err := role.Struct(roleEntity); err != nil {
			continue
		}

		// 获取角色的API列表
		apiIds, err := dao.RoleApi.FindApiIdsByRoleId(ctx, roleEntity.Id)
		if err != nil {
			continue
		}

		if len(apiIds) == 0 {
			continue
		}

		// 获取API详情（只获取启用状态的API）
		apis, err := dao.ApiInfo.FindByApiIds(ctx, apiIds)
		if err != nil {
			continue
		}

		// 为每个启用的API添加casbin策略
		for _, api := range apis {
			// 过滤禁用状态的API
			if api.Status == consts.StatusDisabled {
				continue
			}

			// 添加策略到casbin
			_, err := enforcer.AddPolicy(strconv.FormatInt(roleEntity.Id, 10), api.Path, api.Method)
			if err != nil {
				return gerror.Wrapf(err, "添加casbin策略失败: role=%d, path=%s, method=%s", roleEntity.Id, api.Path, api.Method)
			}
		}
	}

	// 保存策略到数据库
	err = enforcer.SavePolicy()
	if err != nil {
		return gerror.Wrap(err, "保存casbin策略失败")
	}

	return nil
}

// RefreshEnforcer 刷新Casbin Enforcer
func (c *Casbin) RefreshEnforcer(ctx context.Context) error {
	enforcer := global.Casbin()
	if enforcer == nil {
		return gerror.New("Casbin enforcer未初始化")
	}

	// 重新加载策略
	err := enforcer.LoadPolicy()
	if err != nil {
		return gerror.Wrap(err, "重新加载casbin策略失败")
	}

	return nil
}

// GetPolicies 获取Casbin策略列表
func (c *Casbin) GetPolicies(ctx context.Context, roleId *int64) ([]*CasbinPolicyModel, error) {
	enforcer := global.Casbin()
	if enforcer == nil {
		return nil, gerror.New("Casbin enforcer未初始化")
	}

	var rawPolicies [][]string
	var err error

	if roleId != nil {
		// 获取指定角色的策略
		rawPolicies, err = enforcer.GetFilteredPolicy(0, strconv.FormatInt(*roleId, 10))
		if err != nil {
			return nil, gerror.Wrap(err, "获取过滤策略失败")
		}
	} else {
		// 获取所有策略
		rawPolicies, err = enforcer.GetPolicy()
		if err != nil {
			return nil, gerror.Wrap(err, "获取策略失败")
		}
	}

	// 转换为业务模型
	policies := make([]*CasbinPolicyModel, 0, len(rawPolicies))
	for _, policy := range rawPolicies {
		if len(policy) >= 3 {
			policies = append(policies, &CasbinPolicyModel{
				RoleId: policy[0],
				Path:   policy[1],
				Method: policy[2],
			})
		}
	}

	return policies, nil
}

// AddPolicy 添加Casbin策略
func (c *Casbin) AddPolicy(ctx context.Context, roleId, path, method string) error {
	enforcer := global.Casbin()
	if enforcer == nil {
		return gerror.New("Casbin enforcer未初始化")
	}

	added, err := enforcer.AddPolicy(roleId, path, method)
	if err != nil {
		return gerror.Wrap(err, "添加casbin策略失败")
	}

	if !added {
		return gerror.New("策略已存在")
	}

	return nil
}

// RemovePolicy 删除Casbin策略
func (c *Casbin) RemovePolicy(ctx context.Context, roleId, path, method string) error {
	enforcer := global.Casbin()
	if enforcer == nil {
		return gerror.New("Casbin enforcer未初始化")
	}

	removed, err := enforcer.RemovePolicy(roleId, path, method)
	if err != nil {
		return gerror.Wrap(err, "删除casbin策略失败")
	}

	if !removed {
		return gerror.New("策略不存在")
	}

	return nil
}

// GetApisByRole 获取角色对应的API列表（从Casbin中）
func (c *Casbin) GetApisByRole(ctx context.Context, roleId int64) ([]*CasbinApiModel, error) {
	enforcer := global.Casbin()
	if enforcer == nil {
		return nil, gerror.New("Casbin enforcer未初始化")
	}

	policies, err := enforcer.GetFilteredPolicy(0, strconv.FormatInt(roleId, 10))
	if err != nil {
		return nil, gerror.Wrap(err, "获取角色策略失败")
	}

	apis := make([]*CasbinApiModel, 0, len(policies))
	for _, policy := range policies {
		if len(policy) >= 3 {
			apis = append(apis, &CasbinApiModel{
				Path:   policy[1],
				Method: policy[2],
			})
		}
	}

	return apis, nil
}
