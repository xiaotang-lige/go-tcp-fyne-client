package server

import (
	"go_gui/model"
	"go_gui/request"
	"go_gui/tool"
)

type login struct {
}

func (*login) Verify(z, p string) bool {
	data := &model.UserConfig{}
	data.UserId = z
	data.Password = p
	data.Status = 1
	userData, db, err := request.Api.Login.Post(data)
	if err != nil {
		return false
	}
	tool.Api.Token.Save(userData.Token)
	if db.State != 200 {
		return false
	}
	return true
}
