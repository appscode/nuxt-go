package web

import (
  "io/ioutil"
  "net/http"
  "strings"
)

const (
	tplDirname = "templates"
	tplPath    = "/" + tplDirname
	tplPrefix  = tplPath + "/"
)

var (
	StaticFS   http.FileSystem
	TemplateFS http.FileSystem
)

func init() {
	rootFS := RootFS.(vfsgen۰FS)

	staticFS := vfsgen۰FS{}
	for k, v := range rootFS {
		if k == "/" {
			if dir, ok := v.(*vfsgen۰DirInfo); ok {
				for i, e := range dir.entries {
					if e.Name() == tplDirname {
						dir.entries = append(dir.entries[:i], dir.entries[i+1:]...)
					}
				}
			}
			rootFS[k] = v
		}

		match := k == tplPath || strings.HasPrefix(k, tplPrefix)
		if !match {
			staticFS[k] = v
		}
	}
	StaticFS = staticFS

	templateFS := vfsgen۰FS{}
	for k, v := range rootFS {
		if k == tplPath {
			templateFS["/"] = v
		} else if strings.HasPrefix(k, tplPrefix) {
			templateFS[k[len(tplPath):]] = v
		}
	}
	TemplateFS = templateFS
}

func Asset(fs http.FileSystem, name string) ([]byte, error) {
	f, err := fs.Open("/" + name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func AssetNames(fs http.FileSystem) []string {
	realFS := fs.(vfsgen۰FS)
	var results = make([]string, 0, len(realFS))
	for k := range realFS {
		results = append(results, k[1:])
	}
	return results
}
