package request

type api struct {
	Login *login
	Port  *method
}

var Api = new(api)
