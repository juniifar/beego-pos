package impl

import (
	"github.com/beego-pos/internal/app/interface/repository/user"
	"github.com/beego-pos/internal/app/usecase/authorization"
	"github.com/beego-pos/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
)

type AuthorizationImpl struct {
	User user.IUser
}

func New(userRP user.IUser) authorization.IAuthorization {
	return &AuthorizationImpl{User: userRP}
}

func (impl *AuthorizationImpl) Registration(ctx *context.Context, user models.User) error {
	logs.Debug("Test Dulu", 123)

	validate := validation.Validation{}
	validate.Required(user.Email, "emailRequired").Message("email tidak boleh kosong")
	validate.Email(user.Email, "emailValid").Message("email harus valid")

	if validate.HasErrors() {
		for _, err := range validate.Errors {
			logs.Error("error : ", err.Message)
		}
	}

	err := impl.User.SetUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
