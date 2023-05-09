package data

import (
	"fyne.io/fyne/v2/widget"
	"go_gui/model"
)

var Context = make(chan string, 1024)
var MessageData = make(chan *model.Message, 1024)
var MyContext *widget.Entry

func Push() {
	var context string
	for v := range MessageData {
		t := v.CreateTime
		timeLayoutStr := "2006-01-02 15:04:05"
		t.Format(timeLayoutStr)
		xc := timeLayoutStr + "_" + v.Target + ":" + v.Text + "\n"
		context += xc
		MyContext.SetText(context)
	}
}
