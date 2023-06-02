package server

type api struct {
	Login         *login
	Likman        *contacts
	LoginVerify   *verifyLogin
	InformMessage *inform
}

var Api = new(api)
