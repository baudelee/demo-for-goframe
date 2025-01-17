package user

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"

	"demo/api/user/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// return nil, gerror.NewCode(gcode.CodeNotImplemented)
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Name:   req.Name,
		Status: req.Status,
		Age:    req.Age,
	}).WherePri(req.Id).Update()
	return
}
