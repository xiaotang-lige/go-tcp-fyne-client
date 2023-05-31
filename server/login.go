package server

import (
	"go_gui/model"
	"go_gui/request"
	"go_gui/tcp"
	"go_gui/tool"
)

type login struct {
}

func (*login) Verify(z, p string) bool {
	data := &model.UserConfig{}
	data.UserId = z
	data.Password = p
	userData, db, err := request.Api.Login.Post(data)
	if err != nil {
		return false
	}
	if db.State != 200 {
		return false
	}
	tool.Api.Token.Save(userData.Token)
	tool.Api.Config.Save(userData.UserId)
	return true
}
func (*login) StartTcp() bool {
	b := Api.LoginVerify.LoginConfig()
	tcp.TcpDialListen()
	tcp.Main()
	go MessageShow(model.ClientUserConfig.UserId)
	return b
}
