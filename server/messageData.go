package server

import (
	"go_gui/model"
	"go_gui/tcp"
)

func MessageShow(userId string) {
	var context string
	for v := range model.MessageDataPut {
		tcp.TcpServerApi.MessageWrite(tcp.TcpDb, v)
		v.UserId = userId
		t := v.CreateTime
		timeLayoutStr := "2006-01-02 15:04:05"
		t.Format(timeLayoutStr)
		xc := timeLayoutStr + "_" + v.Target + ":" + v.Text + "\n"
		context += xc
		model.MessageDataShow.SetText(context)
	}
}
