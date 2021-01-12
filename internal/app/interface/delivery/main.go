package delivery

//type MainController struct {
//	web.Controller
//}
//
//func (mc *MainController) Get() {
//	mc.Ctx.WriteString("OK")
//}

func (ac *Delivery) Index(){
	ac.Ctx.WriteString("OK")
}