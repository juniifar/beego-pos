package delivery

import (
	"github.com/beego-pos/models"
	"github.com/beego/beego/v2/core/logs"
)

func (ac *Delivery) bindRegister(user *models.User) {
	context := ac.Ctx

	err := context.Input.Bind(&user.Email, "email")
	if err != nil {
		logs.Error("Error : ", err)
		return
	}

	err = context.Input.Bind(&user.Password, "password")
	if err != nil {
		logs.Error("Error : ", err)
		return
	}

	err = context.Input.Bind(&user.Repassword, "repassword")
	if err != nil {
		logs.Error("Error : ", err)
		return
	}
}

func (ac *Delivery) Register() {
	logs.Debug("my book is bought in the year of ", 2016)

	var user models.User
	ac.bindRegister(&user)

	s := ac.Ctx

	err := ac.Authorization.Registration(s, user)
	if err != nil {
		logs.Error("Error : ", err)
	}
	ac.Ctx.WriteString("pong")
}
