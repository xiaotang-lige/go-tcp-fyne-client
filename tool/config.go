package tool

import "os"

const userPath = "/config/user.txt"

type user struct {
}

func (*user) Save(token string) string {
	File, err := os.OpenFile(Api.Path.Project()+userPath, os.O_TRUNC|os.O_CREATE, 0666)
	defer File.Close()
	if err != nil {
		return ""
	}
	_, err = File.WriteString(token)
	if err != nil {
		return ""
	}
	return token
}
func (*user) Get() string {
	file, err := os.OpenFile(Api.Path.Project()+userPath, os.O_RDONLY, 0666)
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
