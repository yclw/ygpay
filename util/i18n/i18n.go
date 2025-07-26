package i18n

var (
	UseLang = map[string]string{}
)

func AddUseLang(lang string) bool {
	// 判断是否支持存在该语言
	if !IsLang(lang) {
		return false
	}
	// 添加到使用语言列表
	UseLang[lang] = lang
	return true
}

func IsUseLang(lang string) bool {
	// 判断是否在使用语言列表中
	_, ok := UseLang[lang]
	return ok
}

func IsLang(lang string) bool {
	_, ok := LangMap[lang]
	return ok
}
