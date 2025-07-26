package token

// TODO：自动刷新，多端登录相关

import (
	"context"
	"fmt"

	"yclw/ygpay/util/token"

	"github.com/golang-jwt/jwt/v5"
)

const (
	CacheRefreshTokenKey = "refresh_token_"
)

// RefreshIdentity 刷新身份
type RefreshIdentity struct {
	Uid string `json:"uid"              dc:"用户ID"`
}

// RefreshClaims 刷新令牌声明
type RefreshClaims struct {
	Identity RefreshIdentity `json:"Identity"`
	jwt.RegisteredClaims
}

var RefreshJwt *RefreshJwtHandler

type RefreshJwtHandler struct {
	*token.JwtHandler
}

func NewRefreshJwtHandler(c *token.TokenConfig, method jwt.SigningMethod) *RefreshJwtHandler {
	return &RefreshJwtHandler{
		token.NewJwtHandler(c, method, &RefreshClaims{}),
	}
}

// CreateToken 创建jwt token
func (h *RefreshJwtHandler) CreateToken(ctx context.Context, identity RefreshIdentity) (header string, expires int64, err error) {
	claims := &RefreshClaims{
		Identity: identity,
	}
	header, expires, err = h.JwtHandler.CreateToken(ctx, claims)
	return
}

func (h *RefreshJwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, identity RefreshIdentity, err error) {
	ok, claims, err := h.JwtHandler.VerifyToken(ctx, header)
	if err != nil || !ok {
		return
	}

	// 现在直接进行类型断言即可，因为JWT解析时使用了RefreshClaims
	identity = claims.(*RefreshClaims).Identity

	return
}

// 获得缓存key
func (h *RefreshJwtHandler) GetCacheKey(ctx context.Context, uid string) string {
	return fmt.Sprintf("%s%s", CacheRefreshTokenKey, uid)
}
