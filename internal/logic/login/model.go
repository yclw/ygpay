package login

type LoginModel struct {
	Uid          string `json:"uid"              dc:"用户ID"`
	Username     string `json:"username"        dc:"用户名"`
	Token        string `json:"token"           dc:"登录token"`
	Expires      int64  `json:"expires"         dc:"登录有效期"`
	RefreshToken string `json:"refreshToken"    dc:"刷新token"`
}

type LoginRefreshTokenModel struct {
	Token        string `json:"token"           dc:"登录token"`
	Expires      int64  `json:"expires"         dc:"登录有效期"`
	RefreshToken string `json:"refreshToken"    dc:"刷新token"`
}
