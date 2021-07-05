package protodoc

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/aclisp/protodoc-serv/internal/ierr"
	"github.com/aclisp/protodoc-serv/internal/pbmd"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	pp "github.com/yoheimuta/go-protoparser/v4"
)

type ProtoDoc struct{}

func (p *ProtoDoc) Convert(ctx context.Context, req *ConvertReq, res *ConvertRes) error {
	if req.Proto == "" {
		return ierr.BadRequest("proto file content is empty")
	}

	reader := strings.NewReader(req.Proto)
	parsed, err := pp.Parse(reader,
		pp.WithDebug(false),
		pp.WithPermissive(true),
		pp.WithFilename(filepath.Base(req.Filename)))
	if err != nil {
		return ierr.BadRequest("can not parse proto file: %v", err)
	}

	protoFile := pbmd.ProtoFile{}
	protoFile.ComposeFrom(parsed)
	mark := protoFile.GenerateMarkdown()
	html := markdown.ToHTML([]byte(mark), parser.NewWithExtensions(
		parser.CommonExtensions|parser.AutoHeadingIDs,
	), nil)

	res.Markdown = mark
	res.Html = string(html)
	return nil
}
