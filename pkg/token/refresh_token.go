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

var RefreshJwt *RefreshJwtHandler

type RefreshJwtHandler struct {
	*token.JwtHandler
}

func NewRefreshJwtHandler(c *token.TokenConfig, method jwt.SigningMethod) *RefreshJwtHandler {
	return &RefreshJwtHandler{
		token.NewJwtHandler(c, method),
	}
}

// CreateToken 创建jwt token
func (h *RefreshJwtHandler) CreateToken(ctx context.Context, c RefreshIdentity) (header string, expires int64, err error) {
	header, expires, err = h.JwtHandler.CreateToken(ctx, c)
	return
}

func (h *RefreshJwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, identity RefreshIdentity, err error) {
	ok, refresh, err := h.JwtHandler.VerifyToken(ctx, header)
	if err != nil {
		return
	}
	identity = refresh.(RefreshIdentity)
	return
}

// 获得缓存key
func (h *RefreshJwtHandler) GetCacheKey(ctx context.Context, uid string) string {
	return fmt.Sprintf("%s%s", CacheRefreshTokenKey, uid)
}
