package server

import (
	"go_gui/model"
)

func MessageShow() {
	var context string
	for v := range model.MessageDataPut {
		t := v.CreateTime
		timeLayoutStr := "2006-01-02 15:04:05"
		xc := t.Format(timeLayoutStr) + "_" + v.Target + ":" + v.Text + "\n"
		context += xc
		model.MessageDataShow.SetText(context)
	}
}
