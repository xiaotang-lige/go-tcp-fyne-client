package request

type api struct {
	Login    *login
	Port     *method
	Contacts *contacts
}

var Api = new(api)
