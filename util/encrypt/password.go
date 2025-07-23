package encrypt

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"

	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"
)

const (
	// bcrypt推荐的最小成本因子
	DefaultCost = 12
	// 盐值长度（字节）
	SaltLength = 32
)

// PasswordStrength 密码强度要求
type PasswordStrength struct {
	MinLength     int  // 最小长度
	RequireUpper  bool // 需要大写字母
	RequireLower  bool // 需要小写字母
	RequireDigit  bool // 需要数字
	RequireSymbol bool // 需要特殊字符
}

// DefaultPasswordStrength 默认密码强度要求
var DefaultPasswordStrength = PasswordStrength{
	MinLength:     8,
	RequireUpper:  true,
	RequireLower:  true,
	RequireDigit:  true,
	RequireSymbol: false,
}

// HashPassword 使用bcrypt加密密码
func HashPassword(password string) (hash string, err error) {
	if password == "" {
		return "", gerror.New("密码不能为空")
	}

	// 使用bcrypt生成哈希（自动生成盐值）
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
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

// GenerateSecureToken 生成安全令牌（用于密码重置等）
func GenerateSecureToken(length int) (string, error) {
	if length <= 0 {
		length = SaltLength
	}

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", gerror.Wrap(err, "生成安全令牌失败")
	}

	return hex.EncodeToString(bytes), nil
}

// HashPasswordWithSalt 使用自定义盐值加密密码（兼容老系统）
// 注意：新系统不建议使用，仅用于迁移
func HashPasswordWithSalt(password, salt string) (string, error) {
	// 这里仍使用bcrypt，但加入自定义盐值
	combined := password + salt
	return HashPassword(combined)
}
