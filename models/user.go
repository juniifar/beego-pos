package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id            int64
	Email         string
	Password      string
	Repassword    string
	Lastlogintime time.Time `orm:"null"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}
