package v1

import "github.com/gogf/gf/v2/frame/g"

// LoginConfigReq 获取登录配置
type LoginConfigReq struct {
	g.Meta `path:"/login/config" method:"get" tags:"登录" summary:"获取登录配置"`
}

type LoginConfigRes struct {
	RegisterSwitch bool   `json:"registerSwitch" dc:"注册开关"`
	CaptchaSwitch  bool   `json:"captchaSwitch" dc:"验证码开关"`
	CaptchaType    int    `json:"captchaType" dc:"验证码类型"`
	Protocol       string `json:"protocol" dc:"协议"`
	Policy         string `json:"policy" dc:"政策"`
}

// LoginCaptchaReq 获取登录验证码
type LoginCaptchaReq struct {
	g.Meta `path:"/login/captcha" method:"get" tags:"登录" summary:"获取登录验证码"`
}

type LoginCaptchaRes struct {
	Cid    string `json:"cid" dc:"验证码ID"`
	Base64 string `json:"base64" dc:"验证码"`
}

// 登录统一返回
type LoginRes struct {
	Uid          string `json:"uid"              dc:"用户ID"`
	Username     string `json:"username"        dc:"用户名"`
	Token        string `json:"token"           dc:"登录token"`
	Expires      int64  `json:"expires"         dc:"登录有效期"`
	RefreshToken string `json:"refreshToken"    dc:"刷新token"`
}

// AccountLoginReq 提交账号登录
type AccountLoginReq struct {
	g.Meta   `path:"/login/accountLogin" method:"post" tags:"登录" summary:"账号登录"`
	Username string `json:"username" v:"required|length:1,30" dc:"用户名"`
	Password string `json:"password" v:"required|length:1,30" dc:"密码"`
	Cid      string `json:"cid"  dc:"验证码ID"`
	Code     string `json:"code" dc:"验证码"`
}

type AccountLoginRes struct {
	*LoginRes
}

// LoginLogoutReq 注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/login/logout" method:"post" tags:"登录" summary:"注销登录"`
}

type LoginLogoutRes struct{}

// LoginRefreshTokenReq 刷新token
type LoginRefreshTokenReq struct {
	g.Meta       `path:"/login/refreshToken" method:"post" tags:"登录" summary:"刷新token"`
	RefreshToken string `json:"refreshToken" v:"required" dc:"刷新token"`
}

type LoginRefreshTokenRes struct {
	Token        string `json:"token"           dc:"登录token"`
	Expires      int64  `json:"expires"         dc:"登录有效期"`
	RefreshToken string `json:"refreshToken"    dc:"刷新token"`
}
