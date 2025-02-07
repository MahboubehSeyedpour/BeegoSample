package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id         int       `orm:"auto;pk"`
	Name       string    `orm:"size(100);not null"`
	Email      string    `orm:"size(100);unique"`
	Password   string    `orm:"size(255);not null"`
	IsVerified bool      `orm:"default(false)"`
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}
