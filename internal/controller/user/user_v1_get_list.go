package user

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"

	"demo/api/user/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	// return nil, gerror.NewCode(gcode.CodeNotImplemented)
	res = &v1.GetListRes{}
	err = dao.User.Ctx(ctx).Where(do.User{
		Age:    req.Age,
		Status: req.Status,
	}).Scan(&res.List)
	return
}
