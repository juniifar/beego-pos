package user

import (
	"github.com/beego-pos/models"
	"github.com/beego/beego/v2/server/web/context"
)

type IUser interface {
	SetUser(ctx *context.Context, user models.User) error
}
