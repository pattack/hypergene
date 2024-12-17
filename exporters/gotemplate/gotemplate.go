package gotemplate

import (
	"context"
	"io"
	"text/template"
)

type Exporter struct {
	Context any
}

func (ex Exporter) Export(_ context.Context, src io.Reader, dst io.Writer) error {
	bb, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	t, err := template.New("").Parse(string(bb))
	if err != nil {
		return err
	}

	return t.Execute(dst, ex.Context)
}
