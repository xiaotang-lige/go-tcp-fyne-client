package server

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/boltdb/bolt"
	"go_gui/file"
	"go_gui/tool"
	"log"
	"strconv"
)

type inform struct {
}

var InformLoadMessageInt = make(chan int, 100)
var InforShowPut = make(chan string, 100)
var listenThisUserId = make(chan string, 100)
var linkManLen = make(map[string]int)

// 监听需要发送的页面
func (*inform) Listen() {
	var err error
	file.Db, err = bolt.Open(tool.Api.Path.Project()+"/chatData.db", 0600, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for v := range listenThisUserId {
		InformLoadMessageInt <- linkManLen[v]
	}
}

// 当前页面回调消息
func (*inform) ThisWindow(userId string) {
	listenThisUserId <- userId
}

// 接受消息通知
func (*inform) ShowInform(f *fyne.Container) {
	go func() {
		for i := 0; i < 20; i++ {
			InforShowPut <- "xxxx"
		}
	}()
	for v := range InforShowPut {
		if _, ok := linkManLen[v]; !ok {
			linkManLen[v] = 1
		} else {
			linkManLen[v]++
		}
		f.Add(widget.NewLabel(v + ":" + strconv.Itoa(linkManLen[v]) + "条消息"))
	}

}
