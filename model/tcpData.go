package model

var MessageDataPut = make(chan *Message, 1024)
var MessageDataShow = make(chan *Message, 1024)
