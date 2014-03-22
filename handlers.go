package main

import (
	"encoding/xml"
	// "net/http"
)

var XMLHead string = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

type Response struct {
	Say string `xml:"Say"`
}

const (
	WelcomeText = "Hello. This is facepunch by phone."
)

func newCaller() string {
	Testresponce := Response{Say: WelcomeText}
	output, _ := xml.Marshal(Testresponce)
	return XMLHead + string(output)
}
