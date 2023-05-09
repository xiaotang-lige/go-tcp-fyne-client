package tcp

import (
	"go_gui/model"
	"net"
)

var messagePod = make(chan *model.Message, 1024)
var TcpDb net.Conn
var TcpServerApi = new(Tcp)

type Tcp struct{}

func Main() {
	go TcpServerApi.messageListen(TcpDb)
}
