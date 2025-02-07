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

func (c *CategoryController) Create() {
	var category models.Category
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &category)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request data"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()

	// Check if a category with the same title already exists
	exists := o.QueryTable("category").Filter("Title", category.Title).Exist()
	if exists {
		c.Data["json"] = map[string]string{"error": "Category already exists"}
		c.ServeJSON()
		return
	}

	_, err = o.Insert(&category)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to create category", "details": err.Error()}
	} else {
		c.Data["json"] = category
	}
	c.ServeJSON()
}
