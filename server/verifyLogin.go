package server

import (
	"go_gui/model"
	"go_gui/tool"
)

type verifyLogin struct {
}

func (*verifyLogin) LoginVerify(config *model.UserConfig) {

}
func (*verifyLogin) LoginConfig() (d bool) {
	u := &model.UserConfig{}
	u.UserId = tool.Api.Config.Get()
	u.Token = tool.Api.Token.Get()
	if u.UserId == "" || u.Token == "" {
		return false
	}
	model.ClientUserConfig = u
	return true
}
