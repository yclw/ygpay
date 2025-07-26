package api

import (
	"context"
	"yclw/ygpay/internal/dao"
)

var ApiService = NewApi()

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

// GetOne 获取单个api信息
func (a *Api) GetOne(ctx context.Context, id int64) (res *ApiModel, err error) {
	res = &ApiModel{}

	// 获取api信息
	res.ApiInfo, err = dao.ApiInfo.FindByID(ctx, id)
	if err != nil {
		return
	}

	return
}

// GetAllList 获取所有api列表
func (a *Api) GetAllList(ctx context.Context) (res []*ApiModel, err error) {
	apis, err := dao.ApiInfo.FindAll(ctx)
	if err != nil {
		return
	}

	res = make([]*ApiModel, 0, len(apis))
	for _, api := range apis {
		res = append(res, &ApiModel{ApiInfo: api})
	}
	return
}

// Create 创建api
func (a *Api) Create(ctx context.Context, req *ApiCreateModel) (err error) {
	id, err := dao.ApiInfo.Create(ctx, req.ApiInfo)
	if err != nil {
		return
	}

	req.ApiInfo.Id = id
	return
}

// Update 更新api
func (a *Api) Update(ctx context.Context, req *ApiUpdateModel) (err error) {
	err = dao.ApiInfo.Update(ctx, req.ApiInfo)
	if err != nil {
		return
	}
	return
}

// Delete 删除api
func (a *Api) Delete(ctx context.Context, id int64) (err error) {
	err = dao.ApiInfo.Delete(ctx, id)
	if err != nil {
		return
	}
	return
}
