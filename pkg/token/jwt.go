package token

// TODO：自动刷新，多端登录相关

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

var Jwt *JwtHandler

// jwt 处理器
type JwtHandler struct {
	config *TokenConfig
	cache  *gcache.Cache
}

func NewJwtHandler(c *TokenConfig) *JwtHandler {
	return &JwtHandler{
		config: c,
	}
}

func Init(c *TokenConfig) {
	Jwt = NewJwtHandler(c)
}

// 初始化缓存
func (h *JwtHandler) InitCache(cache *gcache.Cache) {
	h.cache = cache
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
			ExpiresAt: jwt.NewNumericDate(now.Add(gconv.Duration(expires) * time.Second)), // 过期时间
			NotBefore: jwt.NewNumericDate(now.Add(-1000)),                                 // 生效时间
			Issuer:    "ygpay",                                                            // 签发人
			Audience:  jwt.ClaimStrings{user.App},                                         // 受众, 多个受众用逗号分隔
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名
	header, err = token.SignedString([]byte(h.config.SecretKey))
	if err != nil {
		return "", 0, err
	}

	// 缓存token
	if h.config.AutoCache {
		if err = h.CacheToken(ctx, header, user); err != nil {
			return
		}
	}

	return
}

// VerifyToken 验证jwt token
func (h *JwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, identity *Identity, err error) {
	token, err := jwt.ParseWithClaims(header, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.SecretKey), nil
	})
	if ok = token.Valid; err != nil || !ok {
		return
	}

	// 验证自动缓存token
	if h.config.AutoCache {
		if ok, err = h.VerifyCacheToken(ctx, header); err != nil || !ok {
			return
		}
	}

	claims := token.Claims.(*Claims)
	identity = claims.Identity

	return
}

// TODO：

// 缓存token
func (h *JwtHandler) CacheToken(ctx context.Context, header string, identity *Identity) (err error) {
	// TODO：缓存token
	return
}

// 验证缓存token
func (h *JwtHandler) VerifyCacheToken(ctx context.Context, header string) (ok bool, err error) {
	// TODO：验证缓存token
	return
}
