package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Category struct {
	Id        int       `orm:"auto;pk"`
	Title     string    `orm:"size(100);not null;unique"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Category))
}
