package token

import (
	"context"
	"reflect"
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

// jwt 处理器
type JwtHandler struct {
	config     *TokenConfig
	method     jwt.SigningMethod
	claimsType jwt.Claims // Claims类型实例
}

func NewJwtHandler(c *TokenConfig, method jwt.SigningMethod, claimsExample jwt.Claims) *JwtHandler {
	return &JwtHandler{
		config:     c,
		method:     method,
		claimsType: claimsExample,
	}
}

// GetConfig 获取配置
func (h *JwtHandler) GetConfig() *TokenConfig {
	return h.config
}

// CreateToken 创建jwt token
func (h *JwtHandler) CreateToken(ctx context.Context, c jwt.Claims) (header string, expires int64, err error) {
	// 获取当前时间
	now := time.Now()

	// 过期时长(时间戳)
	expires = h.config.Expires

	// 创建RegisteredClaims
	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expires) * time.Second)), // 过期时间
		NotBefore: jwt.NewNumericDate(now.Add(-1000)),                                // 生效时间
		Issuer:    h.config.Issuer,                                                   // 签发人
		Audience:  jwt.ClaimStrings{h.config.Audience},                               // 受众, 多个受众用逗号分隔
	}

	// 使用反射设置RegisteredClaims字段
	if err = setRegisteredClaims(c, registeredClaims); err != nil {
		return "", 0, err
	}

	// 创建token
	token := jwt.NewWithClaims(h.method, c)

	// 签名
	header, err = token.SignedString([]byte(h.config.SecretKey))
	if err != nil {
		return "", 0, err
	}

	return
}

// setRegisteredClaims 使用反射设置RegisteredClaims字段
func setRegisteredClaims(claims jwt.Claims, registeredClaims jwt.RegisteredClaims) error {
	// 获取claims的反射值
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 查找RegisteredClaims字段
	field := v.FieldByName("RegisteredClaims")
	if !field.IsValid() {
		// 如果没有找到RegisteredClaims字段，尝试查找嵌入的jwt.RegisteredClaims
		for i := 0; i < v.NumField(); i++ {
			fieldType := v.Type().Field(i)
			if fieldType.Type == reflect.TypeOf(jwt.RegisteredClaims{}) {
				field = v.Field(i)
				break
			}
		}
	}

	if field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(registeredClaims))
	}

	return nil
}

// VerifyToken 验证jwt token
func (h *JwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, claims jwt.Claims, err error) {
	token, err := jwt.ParseWithClaims(header, h.claimsType, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.SecretKey), nil
	})

	// 先检查错误和token是否为nil
	if err != nil || token == nil {
		return false, nil, err
	}

	// 再检查token是否有效
	if !token.Valid {
		return false, nil, nil
	}

	return true, token.Claims, nil
}
