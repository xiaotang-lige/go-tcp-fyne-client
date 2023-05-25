package model

type Response struct {
	Len   int         `json:"len"`
	State int         `json:"state"`
	Err   string      `json:"err"`
	Data  interface{} `json:"data"`
}
