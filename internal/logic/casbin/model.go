package casbin

// CasbinApiModel casbin API权限模型
type CasbinApiModel struct {
	Path   string `json:"path"`   // API路径
	Method string `json:"method"` // 请求方法
}

// CasbinPolicyModel casbin策略模型
type CasbinPolicyModel struct {
	RoleId string `json:"roleId"` // 角色ID
	Path   string `json:"path"`   // API路径
	Method string `json:"method"` // 请求方法
}

// CasbinRoleApiModel 角色API权限模型
type CasbinRoleApiModel struct {
	RoleId   int64             `json:"roleId"`   // 角色ID
	RoleName string            `json:"roleName"` // 角色名称
	Apis     []*CasbinApiModel `json:"apis"`     // API列表
}
