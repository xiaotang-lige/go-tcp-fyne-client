package ui

import (
	"encoding/binary"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/boltdb/bolt"
	"go_gui/file"
	"go_gui/server"
	"log"
	"strconv"
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
	//grid := container.NewVBox()
	////go func() {
	////	for  {
	////		time.Sleep(time.Duration(2) * time.Second)
	////		grid.Remove(grid.Objects[0])
	////	}
	////}()
	////var dataIndex map[string]int
	//
	//for i := 0; i < 9; i++ {
	//	grid.Add(widget.NewLabel(strconv.Itoa(i)))
	//}
	grid := widget.NewButton("send", func() {
		t, err := file.Db.Begin(true)
		defer t.Rollback()
		var b *bolt.Bucket
		if b = t.Bucket([]byte("Kathryn Scott")); b == nil {
			b, err = t.CreateBucket([]byte("Kathryn Scott"))
		}
		//x := b.Cursor()
		//x.Last()
		//ba := x.Bucket()
		id, _ := b.NextSequence()

		err = b.Put(itob(int(id)), []byte(strconv.FormatUint(id, 10)))
		//err = b.Put([]byte("1"), []byte("Kathryn Scott"))
		if err != nil {
			log.Println(err)
			t.Rollback()
			return
		}
		t.Commit()
		server.InforShowPut <- "Kathryn Scott"
		server.InformLoadMessageInt <- 1
	})
	return container.NewBorder(nil, nil, nil, nil, grid)
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
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
