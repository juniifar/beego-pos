package authorization

import (
	"github.com/beego-pos/models"
	"github.com/beego/beego/v2/server/web/context"
)

// IAuthorization interface
type IAuthorization interface {
	Registration(ctx *context.Context, user models.User) error
}
