package member

import (
	"context"

	v1 "yclw/ygpay/api/member/v1"
	"yclw/ygpay/internal/logic/member"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	// 构建筛选参数
	filter := &member.MemberListFilter{
		Username:  req.Username,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Mobile:    req.Mobile,
		RoleId:    req.RoleId,
		Sex:       req.Sex,
		Status:    req.Status,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		SortField: req.SortField,
		SortDesc:  req.SortDesc,
	}

	// 调用带筛选的查询方法
	models, total, err := c.MemberService.GetListWithFilter(ctx, req.Page, req.Size, filter)
	if err != nil {
		return
	}

	// 构建响应
	res = &v1.GetListRes{
		Total: total,
	}
	res.List = make([]*v1.MemberModel, 0, len(models))
	for _, model := range models {
		res.List = append(res.List, c.memberModelToV1(model))
	}
	return
}
