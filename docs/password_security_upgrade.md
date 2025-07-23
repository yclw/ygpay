# 密码安全升级指南

## 🚨 当前问题分析

### 安全隐患
1. **MD5哈希不安全** - 易受彩虹表攻击
2. **简单盐值处理** - 16字符盐值长度不够
3. **AES ECB模式** - 存在模式攻击风险
4. **缺乏密码策略** - 无复杂度要求和历史检查

## 🏢 企业级最佳实践

### 1. 密码哈希算法
- ✅ **bcrypt** (推荐) - 自带盐值，计算复杂度可调
- ✅ **scrypt** - 内存困难函数，抗ASIC攻击
- ✅ **Argon2** - 最新标准，抗时间和内存攻击
- ❌ **MD5/SHA1** - 已不安全
- ❌ **简单SHA256** - 需要足够的盐值和迭代次数

### 2. 密码策略
```go
type PasswordPolicy struct {
    MinLength     int  // 最小8位
    RequireUpper  bool // 需要大写字母
    RequireLower  bool // 需要小写字母
    RequireDigit  bool // 需要数字
    RequireSymbol bool // 需要特殊字符（可选）
    MaxAgeDays    int  // 密码有效期（如90天）
    HistoryCount  int  // 不能重复使用的历史密码数量
}
```

## 🔧 升级实施步骤

### 1. 添加依赖
```bash
go get golang.org/x/crypto/bcrypt
```

### 2. 数据库升级
```sql
-- 执行 manifest/sql/upgrade_password_security.sql
-- ⚠️ 升级前请备份数据库！

-- 主要变更：
-- password_hash: char(32) → varchar(100)  支持bcrypt
-- salt: char(16) → varchar(32) NULL      兼容旧版本
-- 新增: password_reset_token, password_reset_at
-- 新增: 密码策略表、密码历史表、登录尝试表
```

### 3. 代码升级

#### 新的密码处理
```go
// 创建密码（注册）
err := security.PwdManager.CreatePassword(ctx, memberID, plainPassword)

// 修改密码
err := security.PwdManager.ChangePassword(ctx, memberID, oldPassword, newPassword)

// 密码重置申请
token, err := security.PwdManager.ResetPasswordRequest(ctx, email)

// 密码重置确认
err := security.PwdManager.ResetPasswordConfirm(ctx, token, newPassword)
```

#### 向后兼容
```go
// 登录时自动检测密码类型
func checkPassword(input, salt, hash string) error {
    // 1. 优先尝试bcrypt验证
    if isBcryptHash(hash) {
        return encrypt.VerifyPassword(plainPassword, hash)
    }
    
    // 2. 兼容旧的MD5方式
    if salt != "" && len(hash) == 32 {
        // 旧密码验证逻辑
    }
    
    return errors.New("密码错误")
}
```

## 🛡️ 安全特性

### 1. 密码强度验证
```go
strength := encrypt.PasswordStrength{
    MinLength:     8,
    RequireUpper:  true,
    RequireLower:  true,
    RequireDigit:  true,
    RequireSymbol: false, // 企业可选
}
```

### 2. 防暴力破解
- 登录尝试次数限制
- IP地址锁定
- 验证码机制
- 账户锁定策略

### 3. 密码重置安全
- 安全令牌生成（32字节随机）
- 令牌过期时间（24小时）
- 一次性使用令牌

### 4. 传输安全
- 建议使用AES-GCM模式替代ECB
- 考虑使用HTTPS加密传输
- 前端密码强度检测

## 📋 迁移计划

### 阶段1：基础设施准备
- [ ] 执行数据库迁移脚本
- [ ] 部署新的密码处理代码
- [ ] 更新密码策略配置

### 阶段2：向后兼容
- [ ] 现有用户仍可正常登录（MD5验证）
- [ ] 用户修改密码时自动升级为bcrypt
- [ ] 监控新旧密码类型比例

### 阶段3：全面升级
- [ ] 通知用户更新密码获得更好安全性
- [ ] 强制要求重要账户升级密码
- [ ] 逐步停止支持MD5验证

### 阶段4：安全增强
- [ ] 启用密码历史检查
- [ ] 实施密码过期策略
- [ ] 添加登录异常检测

## ⚠️ 注意事项

### 1. 兼容性
- 现有用户密码仍可正常验证
- 升级为渐进式，不影响现有业务
- 建议用户主动更改密码

### 2. 性能考虑
- bcrypt计算相对耗时（安全与性能平衡）
- 可调整cost参数（默认12）
- 考虑缓存策略

### 3. 安全建议
- 定期更新密码策略
- 监控异常登录行为
- 实施多因素认证（推荐）
- 定期安全审计

## 🔗 相关文件

- `util/encrypt/password.go` - 密码加密工具
- `internal/logic/security/password.go` - 密码管理服务
- `internal/logic/login/login.go` - 登录验证逻辑
- `manifest/sql/upgrade_password_security.sql` - 数据库升级脚本

## 🏆 企业标准对比

| 功能 | 当前实现 | 企业标准 | 升级后 |
|-----|---------|----------|--------|
| 哈希算法 | MD5 | bcrypt/Argon2 | ✅ bcrypt |
| 盐值长度 | 16字符 | 32字节+ | ✅ bcrypt内置 |
| 密码策略 | 无 | 复杂度要求 | ✅ 可配置 |
| 历史检查 | 无 | 防重复使用 | ✅ 已实现框架 |
| 重置安全 | 无 | 安全令牌 | ✅ 32字节令牌 |
| 传输加密 | AES-ECB | AES-GCM/HTTPS | 🔄 建议升级 |

通过这次升级，你的密码安全将达到企业级标准，大大提升系统的整体安全性。 