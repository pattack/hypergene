package gotemplate_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pattack/hypergene/exporters/gotemplate"
)

func TestExporter_Export(t *testing.T) {
	exporter := gotemplate.Exporter{}

	in := bytes.NewBufferString(`{{ .value }}`)
	out := new(strings.Builder)
	exporter.Context = map[string]string{
		"value": "Pouyan",
	}
	err := exporter.Export(context.Background(), in, out)
	require.NoError(t, err)
	assert.Equal(t, `Pouyan`, out.String())

	in.Reset()
	in.WriteString(`{{ .Value }}`)
	out.Reset()
	exporter.Context = struct {
		Value string
	}{
		Value: "Inka",
	}
	err = exporter.Export(context.Background(), in, out)
	require.NoError(t, err)
	assert.Equal(t, `Inka`, out.String())
}
