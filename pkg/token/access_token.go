package token

import (
	"context"
	"yclw/ygpay/util/token"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
)

// AccessIdentity 访问身份
type AccessIdentity struct {
	Uid      string      `json:"uid"              description:"用户ID"`
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

var AccessJwt *AccessJwtHandler

type AccessJwtHandler struct {
	*token.JwtHandler
}

func NewAccessJwtHandler(c *token.TokenConfig, method jwt.SigningMethod) *AccessJwtHandler {
	return &AccessJwtHandler{
		token.NewJwtHandler(c, method),
	}
}

// CreateToken 创建jwt token
func (h *AccessJwtHandler) CreateToken(ctx context.Context, c AccessIdentity) (header string, expires int64, err error) {
	header, expires, err = h.JwtHandler.CreateToken(ctx, c)
	return
}

func (h *AccessJwtHandler) VerifyToken(ctx context.Context, header string) (ok bool, identity AccessIdentity, err error) {
	ok, refresh, err := h.JwtHandler.VerifyToken(ctx, header)
	if err != nil {
		return
	}
	identity = refresh.(AccessIdentity)
	return
}
