package simple

// 简单工具

import (
	"context"

	"yclw/ygpay/util/encrypt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

// TODO: 需要修改

// ConfigMaskDemoField 演示环境下需要隐藏的配置
var ConfigMaskDemoField = map[string]struct{}{
	// 邮箱
	"smtpUser": {}, "smtpPass": {},

	// 云存储
	"uploadUCloudPublicKey": {}, "uploadUCloudPrivateKey": {}, "uploadCosSecretId": {}, "uploadCosSecretKey": {},
	"uploadOssSecretId": {}, "uploadOssSecretKey": {}, "uploadQiNiuAccessKey": {}, "uploadQiNiuSecretKey": {},

	// 地图
	"geoAmapWebKey": {},

	// 短信
	"smsAliYunAccessKeyID": {}, "smsAliYunAccessKeySecret": {}, "smsTencentSecretId": {}, "smsTencentSecretKey": {},

	// 支付
	"payWxPayMchId": {}, "payWxPaySerialNo": {}, "payWxPayAPIv3Key": {}, "payWxPayPrivateKey": {}, "payQQPayMchId": {}, "payQQPayApiKey": {},

	// 微信
	"officialAccountAppSecret": {}, "officialAccountToken": {}, "officialAccountEncodingAESKey": {}, "openPlatformAppSecret": {},
	"openPlatformToken": {}, "openPlatformEncodingAESKey": {},
}

var RequestEncryptKey = []byte("f080a463654b2279")

// RouterPrefix 获取应用路由前缀
func RouterPrefix(ctx context.Context, app string) string {
	return g.Cfg().MustGet(ctx, "router."+app+".prefix", "/"+app+"").String()
}

// DefaultErrorTplContent 获取默认的错误模板内容
func DefaultErrorTplContent(ctx context.Context) string {
	return gfile.GetContents(g.Cfg().MustGet(ctx, "viewer.paths").String() + "/error/default.html")
}

// CheckPassword 检查密码
func CheckPassword(input, salt, hash string) (err error) {
	// 解码密码
	password, err := gbase64.Decode([]byte(input))
	if err != nil {
		err = gerror.New("密码解析失败")
		return
	}
	// 解密密码
	password, err = encrypt.AesECBDecrypt(password, RequestEncryptKey)
	if err != nil {
		err = gerror.New("密码解析失败")
		return
	}

	// 验证密码
	if hash != gmd5.MustEncryptString(string(password)+salt) {
		err = gerror.New("用户名或密码错误")
		return
	}
	return
}

// SafeGo 安全的调用协程，遇到错误时输出错误日志而不是抛出panic
func SafeGo(ctx context.Context, f func(ctx context.Context), lv ...int) {
	g.Go(ctx, f, func(ctx context.Context, err error) {
		var level = glog.LEVEL_ERRO
		if len(lv) > 0 {
			level = lv[0]
		}
		Logf(level, ctx, "SafeGo exec failed:%+v", err)
	})
}

func Logf(level int, ctx context.Context, format string, v ...interface{}) {
	switch level {
	case glog.LEVEL_DEBU:
		g.Log().Debugf(ctx, format, v...)
	case glog.LEVEL_INFO:
		g.Log().Infof(ctx, format, v...)
	case glog.LEVEL_NOTI:
		g.Log().Noticef(ctx, format, v...)
	case glog.LEVEL_WARN:
		g.Log().Warningf(ctx, format, v...)
	case glog.LEVEL_ERRO:
		g.Log().Errorf(ctx, format, v...)
	case glog.LEVEL_CRIT:
		g.Log().Criticalf(ctx, format, v...)
	case glog.LEVEL_PANI:
		g.Log().Panicf(ctx, format, v...)
	case glog.LEVEL_FATA:
		g.Log().Fatalf(ctx, format, v...)
	default:
		g.Log().Errorf(ctx, format, v...)
	}
}
