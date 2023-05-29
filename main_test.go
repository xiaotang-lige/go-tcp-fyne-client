package main

import (
	"go_gui/server"
	"log"
	"testing"
)

func Test(t *testing.T) {
	data, _ := server.Api.Likman.ShowAll()
	for _, v := range data {
		log.Println(v)
	}
}
