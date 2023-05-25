package request

import (
	"bytes"
	"encoding/json"
	"go_gui/tool"
	"io"
	"log"
	"net/http"
	"os"
)

const url = "http://127.0.0.1:8080"
const tokenPath = "/request/token.txt"

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
	d, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest(t.m, url+path, bytes.NewBuffer(d))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", queryToken())
	clinet := &http.Client{}
	clientDb, err := clinet.Do(req)
	defer clientDb.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(clientDb.Body)
	saveToken(clientDb)
	return body, err
}
func saveToken(r *http.Response) string {
	token := r.Header.Get("Token")
	tokenFile, err := os.OpenFile(tool.ProjectPath()+tokenPath, os.O_TRUNC|os.O_CREATE, 0666)
	defer tokenFile.Close()
	if err != nil {
		return ""
	}
	_, err = tokenFile.WriteString(token)
	if err != nil {
		return ""
	}
	return token
}
func queryToken() string {
	file, err := os.OpenFile(tool.ProjectPath()+tokenPath, os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return ""
	}
	token := make([]byte, 100)
	i, err := file.Read(token)
	if err != nil {
		return ""
	}
	return string(token[:i])
}
