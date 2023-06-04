package server

import (
	"go_gui/model"
	"go_gui/tcp"
	"go_gui/tool"
)

type initmy struct {
}

func (*initmy) Start() bool {

	id := tool.Api.Config.Get()
	token := tool.Api.Token.Get()
	if id == "" || token == "" {
		return false
	}
	model.ConstMy = &model.ThisUserConfig{
		Id:    id,
		Token: token,
	}
	tcp.TcpDialListen()
	tcp.Main()
	go MessageShow()
	go Api.InformMessage.Listen()
	return true
}
