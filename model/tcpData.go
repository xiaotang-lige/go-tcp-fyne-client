package model

import (
	"fyne.io/fyne/v2/widget"
)

var MessageDataPull = make(chan *Message, 1024)
var MessageDataPut = make(chan *Message, 1024)
var MessageDataShow *widget.Entry
