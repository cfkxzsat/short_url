package controllers

import (
	"net/url"

	"github.com/cfkxzsat/short_url/utils"

	"github.com/cfkxzsat/short_url/models"
)

type UrlController struct {
	BaseController
}

func (c *UrlController) Get() {
	// c.Ctx.WriteString("This is UrlController:Get")
	sourceUrl := c.Ctx.Input.Query("sourceUrl")
	// fmt.Println(sourceUrl)
	sourceUrl = "http://" + sourceUrl
	_, err := url.ParseRequestURI(sourceUrl)
	if err != nil {
		c.error("Url is not a valid url:" + err.Error())
	}

	u := &models.Url{
		SourceUrl: sourceUrl,
	}
	u.GenId()
	u.ShortUrl = utils.IdToString(u.Id)
	u.Save()
	c.success(u)
}
