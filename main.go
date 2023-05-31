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
	"go_gui/ui"
	"strconv"
	"time"
)

func main() {
	//os.Unsetenv("FYNE_FONT")

	ui.Main()
}
func uiMain() {
	myApp := app.New()
	w := myApp.NewWindow("miniChat")
	var data []string
	for i := 0; i < 99; i++ {
		data = append(data, strconv.Itoa(i))
	}
	content := container.NewMax()
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])

		})
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Wrapping = fyne.TextWrapWord
	context := container.NewVBox()
	list.OnSelected = func(id widget.ListItemID) {
		context.RemoveAll()
		context.Add(MessageShow(data[id]))
	}
	tutorial := container.NewBorder(context, nil, nil, nil, content)
	split := container.NewHSplit(list, tutorial)
	split.Offset = 0.3
	w.SetContent(split)
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
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
