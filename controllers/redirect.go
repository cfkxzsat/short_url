package controllers

type RedirectController struct {
	BaseController
}

func (c *RedirectController) Get() {
	c.Redirect("/api/url", 302)
}
