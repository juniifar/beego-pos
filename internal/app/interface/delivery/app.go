package delivery

import (
	"github.com/beego-pos/internal/app/usecase/authorization"
	"github.com/beego/beego/v2/server/web"
)

type Delivery struct {
	web.Controller
	Authorization authorization.IAuthorization
}

func New(authorization authorization.IAuthorization) *Delivery {
	return &Delivery{Authorization: authorization}
}