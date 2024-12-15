package hypergene

import (
	"context"
	"io"
)

type SchemeReader interface {
	Read(ctx context.Context, address string) (<-chan File, error)
}

type File struct {
	RelPath string
	Content io.ReadCloser
}
