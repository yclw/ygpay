# jwt鉴权 -- pkg

## 相关文件/文件夹：
 - pkg/token/* 鉴权包
 - internal/middleware/jwt.go 鉴权中间件
 - internal/logic/login/* 登录/登出逻辑
  
## 鉴权包

```go
// Claims 令牌声明
type Claims struct {
	*Identity
	jwt.RegisteredClaims
}
```

```go
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
```

```go
// TokenConfig 登录令牌配置
type TokenConfig struct {
	SecretKey       string `json:"secretKey"`       // 密钥
	Expires         int64  `json:"expires"`         // 过期时长(秒)
    AutoCache       bool   `json:"autoCache"`       // 是否自动缓存token
	AutoRefresh     bool   `json:"autoRefresh"`     // 是否自动刷新
	RefreshInterval int64  `json:"refreshInterval"` // 刷新间隔
	MaxRefreshTimes int64  `json:"maxRefreshTimes"` // 最大刷新次数
	MultiLogin      bool   `json:"multiLogin"`      // 是否允许多端登录
}
```

```go
// 全局JwtHandler
var Jwt *JwtHandler
// 初始化全局JwtHandler
func Init(c *TokenConfig)

// 新建jwt handler
func NewJwtHandler(c *TokenConfig) *JwtHandler
```

```go
// CreateToken 创建jwt token
func (h *JwtHandler) CreateToken(ctx context.Context, user *Identity) (header string, expires int64, err error)
// VerifyToken 验证jwt token
func (h *JwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, identity *Identity, err error)
```