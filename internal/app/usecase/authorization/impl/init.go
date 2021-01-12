package impl

import (
	"github.com/beego-pos/internal/app/interface/repository/user"
	"github.com/beego-pos/internal/app/usecase/authorization"
	"github.com/beego-pos/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)


type AuthorizationImpl struct {
	User user.IUser
}

func New(userRP user.IUser) authorization.IAuthorization{
	return &AuthorizationImpl{User: userRP}
}

func (impl *AuthorizationImpl) Registration(ctx *context.Context, user models.User) error {
	logs.Debug("Test Dulu", 123)
	err := impl.User.SetUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
