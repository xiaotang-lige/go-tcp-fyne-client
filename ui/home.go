package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"go_gui/model"
	"go_gui/server"
	"go_gui/tcp"
	"image/color"
	"strconv"
	"time"
)

func newHome(w fyne.Window) {
	messageContext := container.NewMax()
	list := newList(messageContext)
	mes := func(c *fyne.Container, id string) *fyne.Container {
		c.RemoveAll()
		c.Add(newMessageContext(id))
		return c
	}
	mes(messageContext, "1")

	tabs := container.NewAppTabs(
		container.NewTabItem("message", messageList(messageContext)),
		container.NewTabItem("linkman", list),
	)
	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))
	tabs.SetTabLocation(container.TabLocationBottom)

	split := container.NewHSplit(tabs, messageContext)
	split.Offset = 0.3
	w.SetContent(container.NewMax(split))
}
func messageList(c *fyne.Container) *fyne.Container {
	grid := container.NewVBox()
	//go func() {
	//	for  {
	//		time.Sleep(time.Duration(2) * time.Second)
	//		grid.Remove(grid.Objects[0])
	//	}
	//}()
	//var dataIndex map[string]int
	for i := 0; i < 9; i++ {
		grid.Add(widget.NewLabel(strconv.Itoa(i)))
	}
	return container.NewBorder(nil, nil, nil, nil, grid)
}
func newList(c *fyne.Container) *widget.List {
	listData, l := server.Api.Likman.ShowAll()
	list := widget.NewList(
		func() int {
			return l
		},
		func() fyne.CanvasObject {
			hbox := container.NewGridWithColumns(2)
			tible := container.NewMax(widget.NewLabel("table"))
			hbox.Add(tible)
			hbox.Add(widget.NewLabel(""))
			return hbox
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			//o.(*widget.Label).SetText(listData[i].B)
			o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*widget.Label).SetText(listData[i].B)
		})

	list.OnSelected = func(id widget.ListItemID) {
		c.RemoveAll()
		c.Add(newMessageContext(listData[id].B))
	}

	return list
}
func newMessageContext(s string) *fyne.Container {
	table := widget.NewLabel(s)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	text := widget.NewMultiLineEntry()

	//textMax := container.New(layout.NewGridWrapLayout(fyne.NewSize(450, 300)), text)
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	button := widget.NewButton("send", func() {
		sendButton(input, s)
	})
	//

	bottom := container.NewBorder(nil, nil, nil, button, container.NewMax(input))
	con := container.NewMax(text)
	maxContext := container.NewMax(container.NewBorder(table, bottom, nil, nil, con))
	return maxContext
}
func sendButton(input *widget.Entry, target string) {
	model.MessageDataShow <- &model.Message{
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
