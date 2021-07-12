package protodoc

import (
	"context"
)

type Hello struct{}

func (h *Hello) Hello(ctx context.Context, req *HelloReq, res *HelloRes) error {
	res.Data = "hello"
	return nil
}
