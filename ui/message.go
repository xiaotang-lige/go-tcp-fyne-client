package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"go_gui/model"
	"go_gui/server"
	"go_gui/tcp"
	"image/color"
	"time"
)

func newMessageContext(s string) *fyne.Container {
	{
		server.Api.InformMessage.ThisWindow(s) //给消息中心发送回调消息
	}
	table := widget.NewLabel(s)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	model.MessageDataShow = widget.NewMultiLineEntry()

	//textMax := container.New(layout.NewGridWrapLayout(fyne.NewSize(450, 300)), text)
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Wrapping = fyne.TextWrapWord
	button := widget.NewButton("send", func() {
		sendButton(input, s)
	})
	//
	bottom := container.NewBorder(nil, nil, nil, button, container.NewMax(input))
	con := container.NewMax(model.MessageDataShow)
	maxContext := container.NewMax(container.NewBorder(table, bottom, nil, nil, con))

	//var staus []byte
	//func() {
	//	if staus != nil {
	//
	//	}
	//
	//	t, _ := file.Db.Begin(true)
	//	defer t.Rollback()
	//	b := t.Bucket([]byte(s))
	//	if b != nil {
	//		c := b.Cursor()
	//		var ii []string
	//		for k, v := c.First(); k != nil; k, v = c.Next() {
	//			ii = append(ii, string(v))
	//			staus = k
	//		}
	//		log.Println(ii)
	//		t.Rollback()
	//	}
	//	t.Rollback()
	//}()
	//
	//xc1 := func() {
	//	for v := range server.InformLoadMessageInt {
	//		log.Println("监听到了！", v)
	//		t, err := file.Db.Begin(true)
	//		log.Println(err)
	//		b := t.Bucket([]byte(s))
	//		if b != nil {
	//			log.Println("到达了吗")
	//			c := b.Cursor()
	//			for k, v1 := c.Seek(staus); k != nil; k, v1 = c.Next() {
	//				if v == 0 {
	//					break
	//				}
	//				log.Println(string(v1))
	//				staus = k
	//				v--
	//			}
	//		}
	//		t.Rollback()
	//	}
	//}
	//go xc1()

	return maxContext
}

func sendButton(input *widget.Entry, target string) {
	mes := &model.Message{
		UserId:      model.ConstMy.Id,
		Target:      target,
		CreateTime:  time.Time{},
		MessageType: 1,
		Text:        input.Text,
	}
	model.MessageDataPut <- mes
	input.SetText("")
	tcp.TcpServerApi.MessageWrite(tcp.TcpDb, mes)
	//m, _ := json.Marshal(mes)
	//file.Api.Message.Set(target, m)
}
