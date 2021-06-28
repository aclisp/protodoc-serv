package protodoc

import (
	"context"

	"github.com/aclisp/protodoc-serv/internal/json"
	"github.com/micro/go-micro/v2/logger"
)

type ProtoDoc struct{}

func (p *ProtoDoc) Convert(ctx context.Context, req *ConvertReq, res *ConvertRes) error {
	logger.Infof("ProtoDoc.Convert req=%v", json.Stringify(req))
	res.Markdown = "markdown"
	res.Html = "html"
	return nil
}
