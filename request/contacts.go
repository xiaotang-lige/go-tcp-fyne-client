package request

import (
	"encoding/json"
	"go_gui/model"
	"log"
)

type contacts struct {
}

func (*contacts) GetAll() ([]model.Contacts, *model.Response, error) {
	responseDataByte, err := Api.Port.GET().handler(linkmanPath, nil)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	var data []model.Contacts
	response := &model.Response{Data: &data}
	err = json.Unmarshal(responseDataByte, response)
	if response.State != 200 {
		log.Println("返回错误：", response.Err)
	}
	response.Data = ""
	return data, response, err
}
