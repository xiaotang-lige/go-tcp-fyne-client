package request

import (
	"encoding/json"
	"go_gui/model"
)

type login struct {
}

func (*login) Post(data interface{}) (*model.UserConfig, *model.Response, error) {
	responseDataByte, err := Api.Port.POST().handler("/login/query", data)
	if err != nil {
		return nil, nil, err
	}
	userConfig := &model.UserConfig{}
	response := &model.Response{Data: userConfig}
	err = json.Unmarshal(responseDataByte, response)
	return userConfig, response, err
}
