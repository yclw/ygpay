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

	// 获取所有api信息
	apis, err := dao.ApiInfo.FindAll(ctx)
	if err != nil {
		return
	}

	// 转换为ApiModel
	res = make([]*ApiModel, 0, len(apis))
	for _, api := range apis {
		res = append(res, &ApiModel{ApiInfo: api})
	}
	return
}

// GetList 获取api列表
func (a *Api) GetList(ctx context.Context, page, pageSize int) (res []*ApiModel, err error) {

	// 获取所有api信息
	apis, err := dao.ApiInfo.FindWithPage(ctx, page, pageSize)
	if err != nil {
		return
	}

	// 转换为ApiModel
	res = make([]*ApiModel, 0, len(apis))
	for _, api := range apis {
		res = append(res, &ApiModel{ApiInfo: api})
	}
	return
}

// GetListWithFilter 获取api列表（带筛选）
func (a *Api) GetListWithFilter(ctx context.Context, page, pageSize int, filter *ApiListFilter) (res []*ApiModel, total int, err error) {
	// 构建查询选项
	options := a.buildQueryOptions(filter)

	// 获取筛选后的api信息
	apis, total, err := dao.ApiInfo.FindWithPageAndOptions(ctx, page, pageSize, options...)
	if err != nil {
		return
	}

	// 转换为ApiModel
	res = make([]*ApiModel, 0, len(apis))
	for _, api := range apis {
		res = append(res, &ApiModel{ApiInfo: api})
	}
	return
}

// buildQueryOptions 构建查询选项
func (a *Api) buildQueryOptions(filter *ApiListFilter) []dao.QueryOption {
	if filter == nil {
		return nil
	}

	var options []dao.QueryOption

	cols := dao.ApiInfo.Columns()

	// 名称筛选
	if filter.Name != "" {
		options = append(options, dao.WhereLike(cols.Name, "%"+filter.Name+"%"))
	}

	// 路径筛选
	if filter.Path != "" {
		options = append(options, dao.WhereLike(cols.Path, "%"+filter.Path+"%"))
	}

	// 状态筛选
	if filter.Status != nil {
		options = append(options, dao.Where(cols.Status, *filter.Status))
	}

	// 分组筛选
	if filter.GroupName != "" {
		options = append(options, dao.Where(cols.GroupName, filter.GroupName))
	}

	// 请求方法筛选
	if filter.Method != "" {
		options = append(options, dao.Where(cols.Method, filter.Method))
	}

	// 认证筛选
	if filter.NeedAuth != nil {
		options = append(options, dao.Where(cols.NeedAuth, *filter.NeedAuth))
	}

	// 日期范围筛选
	if filter.StartDate != nil || filter.EndDate != nil {
		options = append(options, dao.WhereBetween(cols.CreatedAt, filter.StartDate, filter.EndDate))
	}

	// 自定义排序
	if filter.SortField != "" {
		if filter.SortDesc {
			options = append(options, dao.OrderDesc(filter.SortField))
		} else {
			options = append(options, dao.OrderAsc(filter.SortField))
		}
	}

	return options
}

// Create 创建api
func (a *Api) Create(ctx context.Context, req *ApiCreateModel) (err error) {
	// 创建api信息
	_, err = dao.ApiInfo.Create(ctx, req.ApiInfo)
	if err != nil {
		return
	}
	return
}

// Update 更新api
func (a *Api) Update(ctx context.Context, req *ApiUpdateModel) (err error) {
	// 更新api信息
	err = dao.ApiInfo.Update(ctx, req.ApiInfo)
	if err != nil {
		return
	}
	return
}

// Delete 删除api
func (a *Api) Delete(ctx context.Context, id int64) (err error) {
	// 删除api信息
	err = dao.ApiInfo.Delete(ctx, id)
	if err != nil {
		return
	}
	return
}
