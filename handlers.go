package main

import (
	"encoding/xml"
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
	"strconv"
)

var XMLHeader string = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

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
	Testresponce.Say = "."

	InputSetup := Gather{}
	InputSetup.Say = GetSectionsString() + "... Please key what sections you want to browse"
	InputSetup.NumDigi = "1"
	InputSetup.Action = "/sections"
	InputSetup.Method = "GET"

	Testresponce.Gather = InputSetup

	output, _ := xml.Marshal(Testresponce)
	return XMLHeader + string(output)
}

func readSections(rw http.ResponseWriter, req *http.Request) string {
	i, e := strconv.ParseInt(req.URL.Query().Get("Digits"), 10, 64)
	Testresponce := Response{}
	if e != nil || int(i) > len(ListSections()) {
		Testresponce.Say = "I'm sorry that was not a valid responce"
		output, _ := xml.Marshal(Testresponce)
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(output)
	}

	listing, e := GetSectionThreads(ListSections()[int(i)].SID)

	if e != nil {
		Testresponce.Say = "I'm sorry we are unable to get a listing at this time"
		output, _ := xml.Marshal(Testresponce)
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(output)
	}

	HandleCount++
	ThreadCache[HandleCount] = listing

	InputSetup := Gather{}
	InputSetup.Say = "Please key what thread you want to browse"
	InputSetup.NumDigi = "1"
	InputSetup.Action = fmt.Sprintf("/threads/%d", HandleCount)
	InputSetup.Method = "GET"

	Testresponce.Gather = InputSetup

	output := ThreadsText

	for k, v := range listing {
		output += fmt.Sprintf("Press %d for the thread %s... ", k, v.ThreadName)
		debug.Println(v.ThreadName)
	}
	debug.Println(output)
	Testresponce.Say = output
	outputb, e := xml.Marshal(Testresponce)
	if e != nil {
		debug.Println("Oh fuck. ", e)
	}
	return XMLHeader + string(outputb)

}

func GetReturnHandler() Gather {
	ReturnGather := Gather{}
	ReturnGather.Action = "/incoming"
	ReturnGather.Method = "POST"
	ReturnGather.NumDigi = "1"
	ReturnGather.Say = "Press any number to return back to the main page."
	return ReturnGather
}

func readThreadPostNum(rw http.ResponseWriter, req *http.Request, prams martini.Params) string {
	Testresponce := Response{}

	handler, handlerr := strconv.ParseInt(prams["handler"], 10, 64)
	postnum, postnumerr := strconv.ParseInt(prams["postnumber"], 10, 64)

	dnum, dnumerr := strconv.ParseInt(req.URL.Query().Get("Digits"), 10, 64)
	if handlerr != nil || dnumerr != nil || postnumerr != nil {
		Testresponce.Say = "An internal error happened... sorry"
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}
	// Grab the reults that where read out to them

	ReturnGather := Gather{}
	ReturnGather.Action = fmt.Sprintf("/threads/%d/%d", handler, postnum+1)
	ReturnGather.Method = "GET"
	ReturnGather.NumDigi = "1"
	ReturnGather.Say = "Press any number to listen to the next post"
	Testresponce.Gather = ReturnGather

	ThreadList := ThreadCache[int(handler)]
	if ThreadCache == nil {
		Testresponce.Say = "An internal error happened... sorry"
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}

	ThreadPosts, e := GetThreadPosts(ThreadList[dnum].ID)
	if e != nil {
		Testresponce.Say = "An internal error happened... sorry"
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}
	if len(ThreadPosts) < int(postnum) {
		Testresponce.Say = "Oh dear... We are unable to read that thread."
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}
	Testresponce.Say = ThreadPosts[postnum-1].Content
	outputb, e := xml.Marshal(Testresponce)
	if e != nil {
		debug.Println("Oh fuck. ", e)
	}
	return XMLHeader + string(outputb)
}

func readThread(rw http.ResponseWriter, req *http.Request, prams martini.Params) string {
	Testresponce := Response{}
	Testresponce.Gather = GetReturnHandler()

	handler, handlerr := strconv.ParseInt(prams["handler"], 10, 64)

	dnum, dnumerr := strconv.ParseInt(req.URL.Query().Get("Digits"), 10, 64)
	if handlerr != nil || dnumerr != nil {
		Testresponce.Say = "An internal error happened... sorry"
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}
	// Grab the reults that where read out to them

	ReturnGather := Gather{}
	ReturnGather.Action = fmt.Sprintf("/threads/%d/%d", handler, 1)
	ReturnGather.Method = "GET"
	ReturnGather.NumDigi = "1"
	ReturnGather.Say = "Press any number to listen to the next post"
	Testresponce.Gather = ReturnGather

	ThreadList := ThreadCache[int(handler)]
	if ThreadCache == nil {
		Testresponce.Say = "An internal error happened... sorry"
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}

	ThreadPosts, e := GetThreadPosts(ThreadList[dnum].ID)
	if e != nil {
		Testresponce.Say = "An internal error happened... sorry"
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}
	if len(ThreadPosts) < 1 {
		Testresponce.Say = "Oh dear... We are unable to read that thread."
		outputb, e := xml.Marshal(Testresponce)
		if e != nil {
			debug.Println("Oh fuck. ", e)
		}
		Testresponce.Gather = GetReturnHandler()

		return XMLHeader + string(outputb)
	}
	Testresponce.Say = ThreadPosts[0].Content
	outputb, e := xml.Marshal(Testresponce)
	if e != nil {
		debug.Println("Oh fuck. ", e)
	}
	return XMLHeader + string(outputb)
}
