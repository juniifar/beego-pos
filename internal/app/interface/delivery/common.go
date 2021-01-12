package delivery

func (ac *Delivery) Ping() {
	ac.Ctx.WriteString("pong")
}
