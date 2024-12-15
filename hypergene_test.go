package hypergene_test

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pattack/hypergene"
	"github.com/pattack/hypergene/exporters/dummy"
)

func TestExport(t *testing.T) {
	var hg hypergene.HyperGene

	wdir, err := os.MkdirTemp("", "hypergene-test-")
	require.NoError(t, err, "unable to create working dir")
	defer func() {
		_ = os.RemoveAll(wdir)
	}()

	schDir, err := os.MkdirTemp(wdir, "sch-")
	require.NoError(t, err, "unable to create scheme dir")

	// fill in scheme files

	err = os.MkdirAll(filepath.Join(schDir, "cmd", "pouyan"), 0o700)
	require.NoError(t, err, "unable to create scheme dir")

	require.NoError(t, os.WriteFile(
		filepath.Join(schDir, "cmd", "pouyan", "main.go.tpl"),
		[]byte(`package main

func main() {
  println("Hello {{ .v.map.Title }}")
}
`),
		0o700,
	))

	outDir, err := os.MkdirTemp(wdir, "out-")
	require.NoError(t, err, "unable to output dir")

	err = hg.Export(context.Background(), dummy.Exporter{}, schDir, outDir)
	require.NoError(t, err)

	dee := make([]string, 0, 10)
	outFiles := make(map[string]fs.DirEntry)
	err = filepath.WalkDir(outDir, func(path string, d fs.DirEntry, serr error) error {
		relPath, _ := filepath.Rel(outDir, path)
		dee = append(dee, relPath)
		outFiles[path] = d

		return nil
	})
	require.NoError(t, err)

	assert.ElementsMatch(t, dee, []string{
		".",
		"cmd",
		filepath.FromSlash("cmd/pouyan"),
		filepath.FromSlash("cmd/pouyan/main.go"),
	})

	bb, err := os.ReadFile(filepath.Join(outDir, filepath.FromSlash("cmd/pouyan/main.go")))
	require.NoError(t, err)
	assert.Equal(t, string(bb), `package main

func main() {
  println("Hello World!")
}
`)
}
