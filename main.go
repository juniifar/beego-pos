package main

import (
	"github.com/beego-pos/internal/app/interface/delivery"
	userrp "github.com/beego-pos/internal/app/interface/repository/user/impl"
	authuc "github.com/beego-pos/internal/app/usecase/authorization/impl"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"net/http"

	_ "github.com/lib/pq"
)

func page_not_found(rw http.ResponseWriter, r *http.Request){
	_, _ = rw.Write([]byte("Kosong Coy"))
}

func init(){
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil{
		logs.Error("error : ", err)
	}

	connectionString, err  := config.String("dev::pgConnectionString")
	if err != nil{
		logs.Error("Error : ", err)
	}

	logs.Debug("Params : ", connectionString)

	err = orm.RegisterDataBase("default", "postgres", connectionString)
	if err != nil{
		logs.Error("error : ", err)
	}
}

func initDB(){
	// Database alias.
	name := "default"

	// Error.
	err := orm.RunSyncdb(name, false, true)
	if err != nil {
		logs.Error("Error : ", err)
	}
}

func main(){
	initDB()

	orms := orm.NewOrm()

	//define repository
	userRP := userrp.New(orms)

	//define usecase
	authUC := authuc.New(userRP)

	delivery := delivery.New(authUC)

	web.ErrorHandler("404", page_not_found)
	//web.Router("/", &controllers.MainController{})
	web.Router("/", delivery, "get:Index")
	web.Router("/ping", delivery, "get:Ping")
	web.Router("/register", delivery, "post:Register")
	web.Run()
}
