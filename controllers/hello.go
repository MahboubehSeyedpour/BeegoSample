package controllers

import "github.com/beego/beego/v2/server/web"

type HelloController struct {
	web.Controller
}

func (c *HelloController) Get() {
	c.Ctx.WriteString("Hello, Beego!")
}
