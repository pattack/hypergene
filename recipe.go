package hypergene

import (
	"context"
	"io"
)

type Exporter interface {
	Export(ctx context.Context, src io.Reader, dst io.Writer) error
}

type Recipe struct {
	Input   File
	OutPath string
}
