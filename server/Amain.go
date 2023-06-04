package server

type api struct {
	Login         *login
	Likman        *contacts
	Init          *initmy
	InformMessage *inform
}

var Api = new(api)

func Main() {
	Api.Init.Start()
}
