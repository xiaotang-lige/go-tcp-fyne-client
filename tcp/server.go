package tcp

import (
	"bufio"
	"encoding/json"
	"go_gui/file"
	"go_gui/model"
	"io"
	"log"
	"net"
)

func (*Tcp) messageListen(c net.Conn) {
	reader := bufio.NewReader(c)
	for {
		megbyte, err := Decode(reader)
		if err != nil || err == io.EOF {
			log.Println(err)
			return
		}
		msg := &model.Message{}
		json.Unmarshal(megbyte, msg)
		//MyContext.SetText(msg.Text)
		//model.MessageDataPut <- msg
		file.Api.Message.Set(msg.Target, megbyte) //保存消息
		//server.InformLoadMessageInt <- 1
		//model.InforShowPut <- msg.Target //通知标记
		model.MessageDataPut <- msg
	}
}
func (*Tcp) MessageWrite(c net.Conn, mes interface{}) error {
	context, err := json.Marshal(mes)
	if err != nil {
		log.Println(err)
		return err
	}
	by, _ := Encode(context)
	_, err = c.Write(by)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
