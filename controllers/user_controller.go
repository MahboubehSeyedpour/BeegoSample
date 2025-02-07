package controllers

import (
	"beegoSample/models"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	web.Controller
}

func (c *UserController) Register() {

	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request data"}
		c.ServeJSON()
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	o := orm.NewOrm()
	_, err = o.Insert(&user)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to register"}
	} else {
		c.Data["json"] = map[string]string{"message": "User registered successfully"}
	}
	c.ServeJSON()
}

func (c *UserController) Login() {
	var input models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)

	o := orm.NewOrm()
	user := models.User{Email: input.Email}

	err := o.Read(&user, "Email")
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid credentials"}
		c.ServeJSON()
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid credentials"}
	} else {
		c.Data["json"] = map[string]string{"message": "Login successful"}
	}
	c.ServeJSON()
}
