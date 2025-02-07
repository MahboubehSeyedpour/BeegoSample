package controllers

import (
	"beegoSample/models"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type CategoryController struct {
	web.Controller
}

// Create a new category
func (c *CategoryController) Create() {
	var category models.Category
	json.Unmarshal(c.Ctx.Input.RequestBody, &category)

	o := orm.NewOrm()
	_, err := o.Insert(&category)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to create category", "details": err.Error()}
	} else {
		c.Data["json"] = category
	}
	c.ServeJSON()
}
