package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
)

var XMLHead string = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

type Response struct {
	Say    string `xml:"Say"`
	Gather Gather `xml:"Gather,omitempty"`
}

type Gather struct {
	Say     string `xml:"Say"`
	NumDigi string `xml:"numDigits,attr"`
	Action  string `xml:"action,attr"`
	Method  string `xml:"method,attr"`
}

const (
	WelcomeText = "Hello. This is facepunch by phone."
	ThreadsText = "This sections top 10 recently posted threads are as follows"
)

func newCaller() string {
	Testresponce := Response{}
	Testresponce.Say = WelcomeText + GetSectionsString()

	InputSetup := Gather{}
	InputSetup.Say = "Please key what sections you want to browse"
	InputSetup.NumDigi = "1"
	InputSetup.Action = "/sections"
	InputSetup.Method = "GET"

	Testresponce.Gather = InputSetup

	output, _ := xml.Marshal(Testresponce)
	return XMLHead + string(output)
}

func readSections(rw http.ResponseWriter, req *http.Request) string {
	d := req.URL.Query().Get("Digits")
	i, e := strconv.ParseInt(d, 10, 64)
	Testresponce := Response{}
	if e != nil || int(i) > len(ListSections()) {
		Testresponce.Say = "I'm sorry that was not a valid responce"
		output, _ := xml.Marshal(Testresponce)
		return XMLHead + string(output)
	}

	listing, e := GetSectionThreads(ListSections()[int(i)].SID)
	if e != nil {
		Testresponce.Say = "I'm sorry we are unable to get a listing at this time"
		output, _ := xml.Marshal(Testresponce)
		return XMLHead + string(output)
	}

	output := ""

	for k, v := range listing {
		output += fmt.Sprintf("Press %d for the thread %s... ", k, v.ThreadName)
	}

	Testresponce.Say = output
	outputb, _ := xml.Marshal(Testresponce)
	return XMLHead + string(outputb)

}
