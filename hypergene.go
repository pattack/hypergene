package hypergene

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

type HyperGene struct {
	l loader
}

func (hg HyperGene) WithLoaders(ll ...LoadOption) HyperGene {
	for _, l := range ll {
		err := hg.l.add(l)
		if err != nil {
			panic(err)
		}
	}

	return hg
}

func (hg HyperGene) Export(ctx context.Context, ex Exporter, schAddr, outDir string) error {
	ff, err := hg.readLocalDir(ctx, schAddr) // todo: use dynamic SchemeReader: local dir walker, remote archive downloader, ...
	if err != nil {
		return err
	}

	// todo: load values using value loaders

	var wg errgroup.Group
	for inFile := range ff {
		wg.Go(hg.export(ctx, ex, hg.recipe(inFile, outDir)))
	}

	return wg.Wait()
}

func (hg HyperGene) readLocalDir(ctx context.Context, address string) (<-chan File, error) {
	ff := make(chan File)

	go func() {
		defer close(ff)

		_ = filepath.WalkDir(address, func(path string, d fs.DirEntry, err error) error {
			if d.Type().IsDir() {
				return nil
			}

			fh, err := os.Open(path)
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(address, path)
			if err != nil {
				return err
			}

			ff <- File{
				RelPath: relPath,
				Content: fh,
			}

			return nil
		})
	}()

	return ff, nil
}

func (hg HyperGene) recipe(inFile File, outDir string) Recipe {
	return Recipe{
		Input:   inFile,
		OutPath: filepath.Join(append([]string{outDir}, hg.trimExtraExtension(inFile.RelPath)...)...),
	}
}

func (hg HyperGene) trimExtraExtension(path string) []string {
	dir, filename := filepath.Split(path)
	if strings.Count(filename, ".") > 1 {
		filename = filename[:len(filename)-len(filepath.Ext(filename))]
	}

	return []string{dir, filename}
}

func (hg HyperGene) export(ctx context.Context, ex Exporter, rcp Recipe) func() error {
	return func() error {
		err := os.MkdirAll(filepath.Dir(rcp.OutPath), 0o700)
		if err != nil {
			return err
		}

		fh, err := os.Create(rcp.OutPath)
		if err != nil {
			return err
		}
		defer func() {
			_ = rcp.Input.Content.Close()
			_ = fh.Close()
		}()

		return ex.Export(ctx, rcp.Input.Content, fh)
	}
}
