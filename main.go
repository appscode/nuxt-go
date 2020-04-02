package main

import (
	"fmt"
	"net/http"

	"github.com/appscode/nuxt-go/web"
	_ "github.com/shurcooL/vfsgen"
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	// m.Use(macaron.Renderer())
	m.Use(macaron.Static("", macaron.StaticOptions{
		//Prefix:      "",
		//SkipLogging: false,
		//IndexFile:   "",
		//Expires:     nil,
		//ETag:        false,
		FileSystem: web.StaticFS,
	}))

	m.Get("/go", web.HTMLRenderer(), func(ctx *macaron.Context) {
		ctx.Data["HostedOn"] = "GitHub"
		ctx.HTML(200, "go")
	})
	m.Run()
}

func main2() {
	fmt.Println("starting server :3131 ...")
	http.Handle("/", http.FileServer(web.RootFS))
	err := http.ListenAndServe(":3131", nil)
	if err != nil {
		panic(err)
	}
}
