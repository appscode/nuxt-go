// +build ignore

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fsTemplates http.FileSystem = modTimeFS{
		fs: http.Dir("../dist"),
	}
	err := vfsgen.Generate(fsTemplates, vfsgen.Options{
		PackageName:  "web",
		BuildTags:    "",
		VariableName: "RootFS",
		Filename:     "bindata.go",
	})
	if err != nil {
		log.Fatal("%v", err)
	}
}

// modTimeFS is an http.FileSystem wrapper that modifies
// underlying fs such that all of its file mod times are set to the Unix epoch.
type modTimeFS struct {
	fs http.FileSystem
}

func (fs modTimeFS) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return modTimeFile{f}, nil
}

type modTimeFile struct {
	http.File
}

func (f modTimeFile) Stat() (os.FileInfo, error) {
	fi, err := f.File.Stat()
	if err != nil {
		return nil, err
	}
	return modTimeFileInfo{fi}, nil
}

type modTimeFileInfo struct {
	os.FileInfo
}

func (modTimeFileInfo) ModTime() time.Time {
	return time.Unix(0, 0)
}
