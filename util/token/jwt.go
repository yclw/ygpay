package token

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenConfig 登录令牌配置
type TokenConfig struct {
	SecretKey       string `json:"secretKey"`       // 密钥
	Issuer          string `json:"issuer"`          // 签发人
	Audience        string `json:"audience"`        // 受众
	Expires         int64  `json:"expires"`         // 过期时长(秒)
	MaxRefreshTimes int64  `json:"maxRefreshTimes"` // 最大刷新次数（内部没有使用，仅作为可读配置）
	MultiLogin      bool   `json:"multiLogin"`      // 是否允许多端登录（内部没有使用，仅作为可读配置）
}

// Claims 令牌声明
type Claims struct {
	Identity interface{}
	jwt.RegisteredClaims
}

// jwt 处理器
type JwtHandler struct {
	config *TokenConfig
	method jwt.SigningMethod
}

func NewJwtHandler(c *TokenConfig, method jwt.SigningMethod) *JwtHandler {
	return &JwtHandler{
		config: c,
		method: method,
	}
}

// GetConfig 获取配置
func (h *JwtHandler) GetConfig() *TokenConfig {
	return h.config
}

// CreateToken 创建jwt token
func (h *JwtHandler) CreateToken(ctx context.Context, c interface{}) (header string, expires int64, err error) {
	// 获取当前时间
	now := time.Now()

	// 过期时长
	expires = h.config.Expires

	// 创建claims
	claims := Claims{
		c,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expires) * time.Second)), // 过期时间
			NotBefore: jwt.NewNumericDate(now.Add(-1000)),                                // 生效时间
			Issuer:    h.config.Issuer,                                                   // 签发人
			Audience:  jwt.ClaimStrings{h.config.Audience},                               // 受众, 多个受众用逗号分隔
		},
	}

	// 创建token
	token := jwt.NewWithClaims(h.method, claims)

	// 签名
	header, err = token.SignedString([]byte(h.config.SecretKey))
	if err != nil {
		return "", 0, err
	}

	return
}

// VerifyToken 验证jwt token
func (h *JwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, identity interface{}, err error) {
	token, err := jwt.ParseWithClaims(header, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.SecretKey), nil
	})
	if ok = token.Valid; err != nil || !ok {
		return
	}

	claims := token.Claims.(*Claims)
	identity = claims.Identity

	return
}
