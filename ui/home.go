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
	"strconv"
	"time"
)

func newHome(w fyne.Window) {

	messageContext := container.NewMax()
	listConext := container.NewMax(newList(messageContext))
	split := container.NewHSplit(listConext, messageContext)
	split.Offset = 0.3
	w.SetContent(container.NewMax(split))
}

func newList(c *fyne.Container) *widget.List {
	var data []string
	for i := 0; i < 99; i++ {
		data = append(data, strconv.Itoa(i))
	}
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
	list.OnSelected = func(id widget.ListItemID) {
		c.RemoveAll()
		c.Add(newMessageContext(data[id]))
	}
	return list
}
func newMessageContext(s string) *fyne.Container {
	table := widget.NewLabel(s)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	text := widget.NewMultiLineEntry()
	textMax := container.New(layout.NewGridWrapLayout(fyne.NewSize(450, 300)), text)
	input := widget.NewMultiLineEntry()
	button := widget.NewButton("send", func() {

	})
	maxContext := container.NewMax(container.NewVBox(table, textMax, input, button))
	return maxContext
}
func home(w fyne.Window) {
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
	w.SetContent(container.NewVBox(split))
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
