package protodoc

import (
	"context"
)

type ProtoDoc struct{}

func (p *ProtoDoc) Convert(ctx context.Context, req *ConvertReq, res *ConvertRes) error {
	res.Markdown = "markdown"
	res.Html = "<p>html</p>"
	return nil
}
