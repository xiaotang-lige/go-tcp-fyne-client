package server

type api struct {
	Login  *login
	Likman *contacts
}

var Api = new(api)
