package controllers

import (
	"fmt"
	"short-url/models"
)

type UrlController struct {
	BaseController
}

func (c *UrlController) Get() {
	// c.Ctx.WriteString("This is UrlController:Get")
	sourceUrl := c.Ctx.Input.Query("sourceUrl")
	fmt.Println(sourceUrl)

	u := &models.Url{
		SourceUrl: sourceUrl,
	}

	c.success(u)
}
