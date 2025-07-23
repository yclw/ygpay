package token

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
)

// TokenConfig 登录令牌配置
type TokenConfig struct {
	SecretKey       string `json:"secretKey"`       // 密钥
	Expires         int64  `json:"expires"`         // 过期时间
	AutoRefresh     bool   `json:"autoRefresh"`     // 是否自动刷新
	RefreshInterval int64  `json:"refreshInterval"` // 刷新间隔
	MaxRefreshTimes int64  `json:"maxRefreshTimes"` // 最大刷新次数
	MultiLogin      bool   `json:"multiLogin"`      // 是否允许多端登录
}

// Claims 令牌声明
type Claims struct {
	*Identity
	jwt.RegisteredClaims
}

// Identity 身份
type Identity struct {
	Id       int64       `json:"id"              description:"用户ID"`
	Pid      int64       `json:"pid"             description:"上级ID"`
	RoleId   int64       `json:"roleId"          description:"角色ID"`
	RoleKey  string      `json:"roleKey"         description:"角色唯一标识符"`
	Username string      `json:"username"        description:"用户名"`
	RealName string      `json:"realName"        description:"姓名"`
	Avatar   string      `json:"avatar"          description:"头像"`
	Email    string      `json:"email"           description:"邮箱"`
	Mobile   string      `json:"mobile"          description:"手机号码"`
	App      string      `json:"app"             description:"登录应用"`
	LoginAt  *gtime.Time `json:"loginAt"         description:"登录时间"`
}

// cache_prefix 缓存前缀
const (
	CachePrefix     = "token_"     // 登录token
	CachePrefixBind = "token_bind" // 登录用户身份绑定
)

// CacheToken 缓存信息
type CacheToken struct {
	ExpireAt     int64 `json:"exp"` // token过期时间
	RefreshAt    int64 `json:"ra"`  // 刷新时间
	RefreshCount int64 `json:"rc"`  // 刷新次数
}
