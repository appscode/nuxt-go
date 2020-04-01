package main

import (
	"fmt"
	"net/http"

	_ "github.com/shurcooL/vfsgen"
<<<<<<< HEAD
	"github.com/appscode/nuxt-go/web"
=======
	"github.com/tamalsaha/nuxt-go/web"
>>>>>>> 3f926d6... Add macaron server
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
<<<<<<< HEAD
	// m.Use(macaron.Renderer())
=======
	m.Use(macaron.Renderer())
>>>>>>> 3f926d6... Add macaron server
	m.Use(macaron.Static("", macaron.StaticOptions{
		//Prefix:      "",
		//SkipLogging: false,
		//IndexFile:   "",
		//Expires:     nil,
		//ETag:        false,
		FileSystem: web.StaticFS,
	}))

<<<<<<< HEAD
	m.Get("/go", web.HTMLRenderer(), func(ctx *macaron.Context) {
=======
	m.Get("/templates/go", web.HTMLRenderer(), func(ctx *macaron.Context) {
>>>>>>> 3f926d6... Add macaron server
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
