package server

import (
	"go_gui/model"
	"go_gui/request"
)

type login struct {
}

func (*login) Verify(z, p string) bool {
	data := &model.UserConfig{}
	data.UserId = z
	data.Password = p
	_, db, err := request.Api.Login.Post(data)
	if err != nil {
		return false
	}
	if db.State != 200 {
		return false
	}
	return true
}
