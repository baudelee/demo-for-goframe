package user

import (
	"context"
	"demo/internal/dao"

	"demo/api/user/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	// return nil, gerror.NewCode(gcode.CodeNotImplemented)
	res = &v1.GetOneRes{}
	err = dao.User.Ctx(ctx).WherePri(req.Id).Scan(&res.User)
	return
}
