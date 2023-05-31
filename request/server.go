package request

import (
	"bytes"
	"encoding/json"
	"go_gui/model"
	"io"
	"log"
	"net/http"
)

const url = "http://127.0.0.1:8080"

type method struct {
	m string
}

func (t *method) GET() *method {
	return &method{m: "GET"}
}
func (t *method) POST() *method {
	return &method{m: "POST"}
}
func (t *method) PUT() *method {
	return &method{m: "PUT"}
}
func (t *method) DELETE() *method {
	return &method{m: "DELETE"}
}
func (t *method) handler(path string, content interface{}) (response []byte, err error) {
	var data []byte
	if content != nil {
		data, err = json.Marshal(content)
		if err != nil {
			log.Println(err)
		}
	} else {
		data = []byte("")
	}
	req, err := http.NewRequest(t.m, url+path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", model.ClientUserConfig.Token)
	clinet := &http.Client{}
	clientDb, err := clinet.Do(req)
	defer clientDb.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(clientDb.Body)
	return body, err
}
