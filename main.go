package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"go_gui/model"
	"go_gui/tcp"
	"time"
)

// func updateTime(clock *widget.Label) {
//
//	s := ""
//	s =
//	log.Println(s)
//	clock.SetText(s)
//
// }

func main() {
	tcp.TcpDialListen()
	tcp.Main()
	show()

}

func show() {
	myApp := app.New()
	myWindow := myApp.NewWindow("miniChat")
	data.MyContext = widget.NewMultiLineEntry()
	//data.MyContext.Disable()

	go data.Push()
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Wrapping = fyne.TextWrapWord
	button := widget.NewButton("sand", func() {
		mes := &model.Message{
			UserId:      "",
			Target:      "dawang",
			CreateTime:  time.Time{},
			MessageType: 1,
			Text:        input.Text,
			Url:         "",
			File:        nil,
		}
		data.MessageData <- mes
		tcp.TcpServerApi.MessageWrite(tcp.TcpDb, mes)
		input.SetText("")
	})
	context := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 300)), data.MyContext)

	content := container.New(layout.NewVBoxLayout(), context, input, button)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.Size{500, 300})
	myWindow.ShowAndRun()
}
