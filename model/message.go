package model

import "time"

type Message struct {
	UserId      string
	Target      string
	CreateTime  time.Time
	MessageType int32
	Text        string
	Url         string
	File        []byte
}
