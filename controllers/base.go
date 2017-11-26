package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) output(res string) {
	c.Ctx.WriteString(res)
}

func (c *BaseController) success(res interface{}) {
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  res,
	}
	c.ServeJSON()
}

func (c *BaseController) error(errMsg string) {
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"message": errMsg,
	}
	c.ServeJSON()
}
