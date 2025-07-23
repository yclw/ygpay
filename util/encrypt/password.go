package encrypt

import (
	"regexp"

	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"
)

const (
	// bcrypt推荐的最小成本因子
	DefaultCost = 12
)

// PasswordStrength 密码强度要求
type PasswordStrength struct {
	MinLength     int  // 最小长度
	RequireUpper  bool // 需要大写字母
	RequireLower  bool // 需要小写字母
	RequireDigit  bool // 需要数字
	RequireSymbol bool // 需要特殊字符
}

// HashPassword 使用bcrypt加密密码
func HashPassword(password string, cost int) (hash string, err error) {
	if password == "" {
		return "", gerror.New("密码不能为空")
	}

	// 使用bcrypt生成哈希
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", gerror.Wrap(err, "密码加密失败")
	}

	return string(bytes), nil
}

// VerifyPassword 验证密码
func VerifyPassword(password, hash string) error {
	if password == "" || hash == "" {
		return gerror.New("密码或哈希值不能为空")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return gerror.New("密码错误")
		}
		return gerror.Wrap(err, "密码验证失败")
	}

	return nil
}

// ValidatePasswordStrength 验证密码强度
func ValidatePasswordStrength(password string, strength PasswordStrength) error {
	if len(password) < strength.MinLength {
		return gerror.Newf("密码长度不能少于%d位", strength.MinLength)
	}

	if strength.RequireUpper {
		if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
			return gerror.New("密码必须包含大写字母")
		}
	}

	if strength.RequireLower {
		if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
			return gerror.New("密码必须包含小写字母")
		}
	}

	if strength.RequireDigit {
		if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
			return gerror.New("密码必须包含数字")
		}
	}

	if strength.RequireSymbol {
		if matched, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]`, password); !matched {
			return gerror.New("密码必须包含特殊字符")
		}
	}

	return nil
}
