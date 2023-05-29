package server

import (
	"go_gui/model"
	"go_gui/request"
	"log"
)

type contacts struct {
}

func (*contacts) ShowAll() ([]model.Contacts, int) {
	var list []model.Contacts
	list, res, err := request.Api.Contacts.GetAll()
	len := res.Len
	if err != nil {
		log.Println("拉去联系人失败！:", err)
	}
	return list, len

}
