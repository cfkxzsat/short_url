package routers

import (
	"try_beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.RESTRouter("/api/url", &controllers.UrlController{})
	beego.Router("/*", &controllers.RedirectController{})
}
