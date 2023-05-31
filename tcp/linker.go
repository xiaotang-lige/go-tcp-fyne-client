package tcp

import (
	"encoding/json"
	"fmt"
	"go_gui/tool"
	"log"
	"net"
)

func TcpDialListen() error {
	var err error
	TcpDb, err = net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return err
	}
	context, err := json.Marshal(tool.Api.Token.Get())
	if err != nil {
		log.Println(err)
		return err
	}
	by, _ := Encode(context)
	_, err = TcpDb.Write(by)
	if err != nil {
		log.Println("tcp对接服务器失败！:", err)
		return err
	}
	return nil
}
