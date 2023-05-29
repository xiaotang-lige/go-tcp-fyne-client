package tool

import (
	"os"
)

const tokenPath = "/request/token.txt"

type m struct {
}

func (*m) Save(token string) string {
	tokenFile, err := os.OpenFile(Api.Path.Project()+tokenPath, os.O_TRUNC|os.O_CREATE, 0666)
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
func (*m) Get() string {
	file, err := os.OpenFile(Api.Path.Project()+tokenPath, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return ""
	}
	token := make([]byte, 1000)
	i, err := file.Read(token)
	if err != nil {
		return ""
	}
	return string(token[:i])
}
