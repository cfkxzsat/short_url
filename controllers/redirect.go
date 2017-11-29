package controllers

import (
	"strconv"
	"strings"

	"github.com/cfkxzsat/short_url/models"
	"github.com/cfkxzsat/short_url/utils"
)

type RedirectController struct {
	BaseController
}

func (c *RedirectController) Get() {
	urlPath := c.Ctx.Request.URL.Path
	if urlPath == "/" {
		c.TplName = "index.tpl"
		return
	}
	if urlPath == "/favicon.ico" {
		return
	}

	urlPath = strings.Trim(urlPath, "/")

	id := utils.StringToId(urlPath)

	url := &models.Url{
		Id: id,
	}
	err := url.FindById()
	if err != nil {
		c.error("没有找到对应的短连接：" + err.Error() + strconv.Itoa(url.Id) + url.SourceUrl)
		return
	}
	c.Redirect(url.SourceUrl, 302)

}
