package captcha

import (
	"context"

	"github.com/mojocn/base64Captcha"
)

var DefaultCaptcha = NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)

type Captcha struct {
	Driver  base64Captcha.Driver
	Store   base64Captcha.Store
	Captcha *base64Captcha.Captcha
}

func NewCaptcha(driver base64Captcha.Driver, store base64Captcha.Store) *Captcha {
	return &Captcha{
		Driver:  driver,
		Store:   store,
		Captcha: base64Captcha.NewCaptcha(driver, store),
	}
}

// Generate 生成验证码
func (c *Captcha) Generate(ctx context.Context) (id, base64, answer string, err error) {
	id, base64, _, err = c.Captcha.Generate()
	return
}

// Verify 使用内置的验证码存储方式验证输入的验证码是否正确
func (c *Captcha) Verify(id, answer string) bool {
	return c.Store.Verify(id, answer, true)
}
