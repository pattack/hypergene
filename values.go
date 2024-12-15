package hypergene

import (
	"context"
	"io"

	"github.com/janstoon/toolbox/bricks"
)

type Loader interface {
	Load(ctx context.Context, r io.Reader) error
}

type LoadOption struct {
	Extensions []string
	Loader     Loader
}

type loader struct {
}

func (l loader) add(lo LoadOption) error {
	return bricks.ErrUnimplemented
}

func (l loader) loadDir(ctx context.Context, dir string) (any, error) {
	return nil, bricks.ErrUnimplemented
}
