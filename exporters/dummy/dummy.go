package dummy

import (
	"context"
	"io"
)

type Exporter struct{}

func (ex Exporter) Export(ctx context.Context, src io.Reader, dst io.Writer) error {
	_, err := io.Copy(dst, src)

	return err
}
