package v1

import "github.com/gogf/gf/v2/frame/g"

// LoginConfigReq 获取登录配置
type LoginConfigReq struct {
	g.Meta `path:"/login/config" method:"get" tags:"登录" summary:"获取登录配置"`
}

type LoginConfigRes struct {
	RegisterSwitch int    `json:"loginRegisterSwitch" dc:"注册开关"`
	CaptchaSwitch  int    `json:"loginCaptchaSwitch" dc:"验证码开关"`
	CaptchaType    int    `json:"loginCaptchaType" dc:"验证码类型"`
	Avatar         string `json:"loginAvatar" dc:"默认头像"`
	RoleId         int64  `json:"loginRoleId" dc:"默认角色ID"`
	Protocol       string `json:"loginProtocol" dc:"协议"`
	Policy         string `json:"loginPolicy" dc:"政策"`
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
	Id       int64  `json:"id"              dc:"用户ID"`
	Username string `json:"username"        dc:"用户名"`
	Token    string `json:"token"           dc:"登录token"`
	Expires  int64  `json:"expires"         dc:"登录有效期"`
}

// AccountLoginReq 提交账号登录
type AccountLoginReq struct {
	g.Meta   `path:"/login/accountLogin" method:"post" tags:"登录" summary:"账号登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Cid      string `json:"cid"  dc:"验证码ID"`
	Code     string `json:"code" dc:"验证码"`
	IsLock   bool   `json:"isLock"  dc:"是否为锁屏状态"`
}

type AccountLoginRes struct {
	*LoginRes
}

// LoginLogoutReq 注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/login/logout" method:"post" tags:"登录" summary:"注销登录"`
}

type LoginLogoutRes struct{}
