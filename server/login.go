package server

import (
	"go_gui/model"
	"go_gui/request"
	"go_gui/tool"
	"log"
)

type login struct {
}

// 登录后保存账号密码
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
	log.Println(Api.Init.Start())
	return true
}
