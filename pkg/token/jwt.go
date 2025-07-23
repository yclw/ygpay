package token

// TODO：自动刷新，多端登录相关

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

var Jwt *JwtHandler

// jwt 处理器
type JwtHandler struct {
	config *TokenConfig
}

func NewJwtHandler(c *TokenConfig) *JwtHandler {
	return &JwtHandler{
		config: c,
	}
}

func Init(c *TokenConfig) {
	Jwt = NewJwtHandler(c)
}

// CreateToken 创建jwt token
func (h *JwtHandler) CreateToken(ctx context.Context, user *Identity) (header string, expires int64, err error) {
	// 获取当前时间
	now := time.Now()

	// 过期时长
	expires = h.config.Expires

	// 创建claims
	claims := Claims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(gconv.Duration(expires))), // 过期时间
			NotBefore: jwt.NewNumericDate(now.Add(-1000)),                   // 生效时间
			Issuer:    "ygpay",                                              // 签发人
			Audience:  jwt.ClaimStrings{user.App},                           // 受众, 多个受众用逗号分隔
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名
	header, err = token.SignedString([]byte(h.config.SecretKey))
	if err != nil {
		return "", 0, err
	}

	return
}

// VerifyToken 验证jwt token
func (h *JwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, identity *Identity, err error) {
	token, err := jwt.ParseWithClaims(header, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.SecretKey), nil
	})
	if err != nil {
		return
	}
	ok = token.Valid

	claims := token.Claims.(*Claims)
	identity = claims.Identity

	return
}
