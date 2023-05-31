package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"go_gui/model"
	"go_gui/tcp"
	"image/color"
	"time"
)

func newMessageContext(s string) *fyne.Container {
	table := widget.NewLabel(s)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	model.MessageDataShow = widget.NewMultiLineEntry()

	//textMax := container.New(layout.NewGridWrapLayout(fyne.NewSize(450, 300)), text)
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	button := widget.NewButton("send", func() {
		sendButton(input, s)
	})
	//
	bottom := container.NewBorder(nil, nil, nil, button, container.NewMax(input))
	con := container.NewMax(model.MessageDataShow)
	maxContext := container.NewMax(container.NewBorder(table, bottom, nil, nil, con))
	return maxContext
}
func sendButton(input *widget.Entry, target string) {
	model.MessageDataPut <- &model.Message{
		Target:      target,
		CreateTime:  time.Time{},
		MessageType: 1,
		Text:        input.Text,
	}
	input.SetText("")
}

func MessageShow(s string) *fyne.Container {
	la := widget.NewLabel(s)
	data.MyContext = widget.NewMultiLineEntry()
	//data.MyContext.Disable()
	go data.Push()
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Wrapping = fyne.TextWrapWord
	button := widget.NewButton("sand", func() {
		mes := &model.Message{
			Target:      s,
			CreateTime:  time.Now(),
			MessageType: 1,
			Text:        input.Text,
		}
		data.MessageData <- mes
		tcp.TcpServerApi.MessageWrite(tcp.TcpDb, mes)
		input.SetText("")
	})
	dataMycontext := container.New(layout.NewGridWrapLayout(fyne.NewSize(450, 300)), data.MyContext)
	c := container.NewVBox(la, widget.NewSeparator(), dataMycontext, input, button)
	return c
}
