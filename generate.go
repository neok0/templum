package tempel

import (
	"context"
	"errors"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"testing/fstest"
)

func Generate(ctx context.Context, contentPath string, theme Theme, outputPath string) error {
	fsys := os.DirFS(contentPath)

	config, err := config(fsys)
	if err != nil {
		return err
	}

	docs, err := fs.Sub(fsys, "docs")
	if err != nil {
		return err
	}

	pages, err := newPages(docs, ".")
	if err != nil {
		return err
	}

	staticFS, err := staticFS(fsys)
	if err != nil {
		return err
	}

	content := Content{
		Pages:  pages,
		Static: staticFS,
		Config: config,
	}

	return generate(ctx, content, theme, outputPath)
}

func generate(ctx context.Context, content Content, theme Theme, outputPath string) error {
	docFS, err := theme.Render(ctx, content)
	if err != nil {
		return err
	}

	return writeToDisk([]fs.FS{docFS, content.Static}, outputPath)
}

func config(fsys fs.FS) (map[string]string, error) {
	var config map[string]string

	f, err := fsys.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

var ErrUnknownTemplate = errors.New("unknown template")

func docFS(ctx context.Context, fsys fs.FS, theme Theme, config map[string]string) (fs.FS, error) {
	docs, err := fs.Sub(fsys, "docs")
	if err != nil {
		return nil, err
	}

	pages, err := newPages(docs, ".")
	if err != nil {
		return nil, err
	}

	content := Content{
		Pages:  pages,
		Config: config,
	}

	docFS, err := theme.Render(ctx, content)
	if err != nil {
		return nil, err
	}

	return docFS, nil
}

func staticFS(fsys fs.FS) (fs.FS, error) {
	static, err := fs.Sub(fsys, "static")
	if err != nil {
		return nil, err
	}

	out := fstest.MapFS{}

	err = fs.WalkDir(static, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		b, err := fs.ReadFile(static, path)
		if err != nil {
			return err
		}

		out[path] = &fstest.MapFile{Data: b}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}
