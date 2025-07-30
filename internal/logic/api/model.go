package api

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// ApiModel api模型
type ApiModel struct {
	*entity.ApiInfo
}

// ApiUpdateModel api更新模型
type ApiUpdateModel struct {
	Id int64
	*do.ApiInfo
}

// ApiCreateModel api创建模型
type ApiCreateModel struct {
	*do.ApiInfo
}

// ApiListFilter api列表筛选参数
type ApiListFilter struct {
	Name      string      `json:"name"`      // 名称
	Path      string      `json:"path"`      // 路径
	Status    *int        `json:"status"`    // 状态筛选
	GroupName string      `json:"groupName"` // 分组筛选
	Method    string      `json:"method"`    // 请求方法筛选
	NeedAuth  *int        `json:"needAuth"`  // 是否需要认证
	StartDate *gtime.Time `json:"startDate"` // 开始日期
	EndDate   *gtime.Time `json:"endDate"`   // 结束日期
	SortField string      `json:"sortField"` // 排序字段
	SortDesc  bool        `json:"sortDesc"`  // 是否降序
}
