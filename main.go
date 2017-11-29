package main

import (
	"flag"
	_ "try_beego/routers"

	"github.com/astaxie/beego"
)

func main() {
	var port string
	flag.StringVar(&port, "p", ":8080", "port to listen")
	flag.Parse()

	beego.Run(port)
}
