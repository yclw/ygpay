package api

import (
	"yclw/ygpay/internal/model/do"
	"yclw/ygpay/internal/model/entity"
)

type ApiModel struct {
	*entity.ApiInfo
}

type ApiUpdateModel struct {
	Id int64
	*do.ApiInfo
}

type ApiCreateModel struct {
	*do.ApiInfo
}
