package model

import "fyne.io/fyne/v2/widget"

var ConstMy = new(ThisUserConfig) //初始化个人数据
var MessageDataPull = make(chan *Message, 1024)
var MessageDataPut = make(chan *Message, 1024)
var MessageDataShow *widget.Entry
var InformLoadMessageInt = make(chan int, 100)
var InforShowPut = make(chan string, 100)
