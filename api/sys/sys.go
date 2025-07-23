// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys

import (
	"context"

	"yclw/ygpay/api/sys/v1"
)

type ISysV1 interface {
	GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error)
	GetGroupConfig(ctx context.Context, req *v1.GetGroupConfigReq) (res *v1.GetGroupConfigRes, err error)
	UpdateGroupConfig(ctx context.Context, req *v1.UpdateGroupConfigReq) (res *v1.UpdateGroupConfigRes, err error)
	SiteConfig(ctx context.Context, req *v1.SiteConfigReq) (res *v1.SiteConfigRes, err error)
}
