package impl

import (
	"github.com/beego-pos/internal/app/interface/repository/user"
	"github.com/beego-pos/models"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

type UserImpl struct {
	orm orm.Ormer
}

func (impl *UserImpl) SetUser(ctx *context.Context, user models.User) error {
	//o := orm.NewOrm()
	x := new(models.User)
	x.Email = "rubahapi@gmail.com"
	y, err := impl.orm.Insert(x)
	if err != nil{
		logs.Error("Error : ", err)
	}
	logs.Debug("test : ", y)

	return nil
}

func New(orm orm.Ormer) user.IUser{
	return &UserImpl{orm: orm}
}
