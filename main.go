package main

import (
	"flag"
	"short-url/controllers"

	"github.com/astaxie/beego"
)

func main() {
	var port string
	flag.StringVar(&port, "p", ":8080", "port to listen")
	flag.Parse()
	// http.Handle("/", http.FileServer(http.Dir("views")))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	beego.SetStaticPath("/", "views")
	// beego.RESTRouter("/redirect", &controllers.RedirectController{})
	beego.RESTRouter("/api/url", &controllers.UrlController{})
	beego.Run(port)
}
