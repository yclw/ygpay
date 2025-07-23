package url

// 获取url信息

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// UriToMap 将URL参数转为map
func UriToMap(uri string) g.MapStrStr {
	m := make(map[string]string)
	if len(uri) < 1 {
		return nil
	}
	if uri[0:1] == "?" {
		uri = uri[1:]
	}
	pars := strings.Split(uri, "&")
	for _, par := range pars {
		kv := strings.Split(par, "=")
		m[kv[0]] = kv[1]
	}
	return m
}

// MapToUri 将map转为URL参数
func MapToUri(params g.MapStrStr) string {
	escape := ""
	for k, v := range params {
		if escape != "" {
			escape = escape + "&"
		}
		escape = escape + k + "=" + v
	}
	return escape
}

// GetAddr 获取请求中的请求地址，协议+域名/IP:端口
func GetAddr(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}
	var (
		scheme = "http"
		proto  = r.Header.Get("X-Forwarded-Proto")
	)
	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	return fmt.Sprintf(`%s://%s`, scheme, r.Host)
}

// GetDomain 获取请求中的域名，如果请求不是域名则返回空
func GetDomain(ctx context.Context) string {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		g.Log().Info(ctx, "GetDomain ctx not request")
		return ""
	}
	if IsDNSName(r.Host) {
		return r.Host
	}
	return ""
}

// IsDNSName 是否是域名地址
func IsDNSName(s string) bool {
	DNSName := `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	rxDNSName := regexp.MustCompile(DNSName)
	return s != "" && rxDNSName.MatchString(s)
}
