package user

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"

	"demo/api/user/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// return nil, gerror.NewCode(gcode.CodeNotImplemented)
	insertId, err := dao.User.Ctx(ctx).Data(do.User{
		Name:   req.Name,
		Status: v1.StatusOK,
		Age:    req.Age,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{
		Id: insertId,
	}
	return
}
