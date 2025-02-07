package controllers

import (
	"beegoSample/models"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type TaskController struct {
	web.Controller
}

func (c *TaskController) Create() {
	var task models.Task
	json.Unmarshal(c.Ctx.Input.RequestBody, &task)

	o := orm.NewOrm()
	_, err := o.Insert(&task)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to create task"}
	} else {
		c.Data["json"] = task
	}
	c.ServeJSON()
}

func (c *TaskController) GetAll() {
	var tasks []models.Task
	o := orm.NewOrm()
	_, err := o.QueryTable("task").All(&tasks)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to retrieve tasks"}
	} else {
		c.Data["json"] = tasks
	}
	c.ServeJSON()
}

func (c *TaskController) Get() {
	id, _ := c.GetInt(":id")
	task := models.Task{Id: id}

	o := orm.NewOrm()
	err := o.Read(&task)
	if err == orm.ErrNoRows {
		c.Data["json"] = map[string]string{"error": "Task not found"}
	} else {
		c.Data["json"] = task
	}
	c.ServeJSON()
}

func (c *TaskController) Update() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	task := models.Task{Id: id}

	if err := o.Read(&task); err == orm.ErrNoRows {
		c.Data["json"] = map[string]string{"error": "Task not found"}
		c.ServeJSON()
		return
	}

	json.Unmarshal(c.Ctx.Input.RequestBody, &task)
	_, err := o.Update(&task)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to update task"}
	} else {
		c.Data["json"] = map[string]string{"message": "Task updated successfully"}
	}
	c.ServeJSON()
}

func (c *TaskController) Delete() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	_, err := o.Delete(&models.Task{Id: id})

	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to delete task"}
	} else {
		c.Data["json"] = map[string]string{"message": "Task deleted successfully"}
	}
	c.ServeJSON()
}
