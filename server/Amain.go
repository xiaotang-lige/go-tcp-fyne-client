package server

type api struct {
	Login       *login
	Likman      *contacts
	LoginVerify *verifyLogin
}

var Api = new(api)
