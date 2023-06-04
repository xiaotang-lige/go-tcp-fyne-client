package server

import (
	"go_gui/model"
	"go_gui/request"
	"go_gui/tool"
	"log"
)

type contacts struct {
}

func (*contacts) ShowAll() ([]model.Contacts, int) {
	var list []model.Contacts
	list, res, err := request.Api.Contacts.GetAll()
	id := tool.Api.Config.Get()
	for k, v := range list {
		if v.B == id {
			list[k].A, list[k].B = list[k].B, list[k].A
		}
	}
	len := res.Len
	if err != nil {
		log.Println("拉去联系人失败！:", err)
	}
	return list, len

}
