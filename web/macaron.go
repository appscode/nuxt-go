package web

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"gopkg.in/macaron.v1"
)

type templateFileSystem struct {
	files []macaron.TemplateFile
}

func (templates templateFileSystem) ListFiles() []macaron.TemplateFile {
	return templates.files
}

func (templates templateFileSystem) Get(name string) (io.Reader, error) {
	for i := range templates.files {
		if templates.files[i].Name()+templates.files[i].Ext() == name {
			return bytes.NewReader(templates.files[i].Data()), nil
		}
	}
	return nil, fmt.Errorf("file '%s' not found", name)
}

func NewTemplateFileSystem() templateFileSystem {
	fs := templateFileSystem{}
	fs.files = make([]macaron.TemplateFile, 0, 10)

	for _, assetPath := range AssetNames(TemplateFS) {
		var ext string
		if strings.HasSuffix(assetPath, "/index.html") {
			ext = "/index.html"
		} else if strings.HasSuffix(assetPath, ".tmpl") {
			ext = ".tmpl"
		}
		if ext == "" {
			continue
		}

		content, err := Asset(TemplateFS, assetPath)
		if err != nil {
			fmt.Printf("failed to read embedded %s template. %v\n", assetPath, err)
			continue
		}

		fs.files = append(fs.files, macaron.NewTplFile(
			strings.TrimSuffix(assetPath, ext),
			content,
			ext,
		))
	}

	return fs
}

// HTMLRenderer implements the macaron handler for serving HTML templates.
func HTMLRenderer() macaron.Handler {
	return macaron.Renderer(macaron.RenderOptions{
		TemplateFileSystem: NewTemplateFileSystem(),
	})
}
