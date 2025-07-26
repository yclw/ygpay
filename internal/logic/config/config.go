package config

import (
	"context"
	"yclw/ygpay/internal/dao"
	"yclw/ygpay/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var ConfigService = NewConfig()

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

// 获取配置
func (c *Config) GetConfig(ctx context.Context) (res g.Map, err error) {
	confs, err := dao.SysConfig.GetConfig(ctx)
	if err != nil {
		return
	}
	res = g.Map{}
	for _, conf := range confs {
		res[conf.Key] = conf.Value
	}
	return
}

// 根据分组获取配置
func (c *Config) GetGroupConfig(ctx context.Context, group string) (res g.Map, err error) {
	confs, err := dao.SysConfig.GetGroupConfig(ctx, group)
	if err != nil {
		return
	}
	if len(confs) == 0 {
		err = gerror.New("配置分组不存在")
		return
	}
	res = g.Map{}
	for _, conf := range confs {
		res[conf.Key] = conf.Value
	}
	return
}

// 根据key获取配置
func (c *Config) GetConfigByKey(ctx context.Context, key string) (res string, err error) {
	conf, err := dao.SysConfig.GetConfigByKey(ctx, key)
	if err != nil {
		return
	}
	if conf == nil {
		err = gerror.New("配置不存在")
		return
	}
	res = conf.Value
	return
}

// 更新配置
func (c *Config) UpdateConfig(ctx context.Context, key string, value string) (err error) {
	err = dao.SysConfig.Update(ctx, &do.SysConfig{
		Key:   key,
		Value: value,
	})
	return
}
