package main

import (
	"github.com/codegangsta/martini"
	"log"
	"os"
)

var info *log.Logger
var debug *log.Logger
var ThreadCache map[int][]Thread
var HandleCount int

func main() {
	info = log.New(os.Stdout, "[info] ", log.Ltime)
	debug = log.New(os.Stderr, "[debug] ", log.Ltime|log.Lshortfile)
	debug.Println(ThreadCache)
	ThreadCache = make(map[int][]Thread)
	info.Println("Facepunch by phone, S16/03/2014")
	debug.Println("Debug text enabled")

	// Martini uses env vars to set its listening host and port,
	// Here I am setting them. In case you where wondering 3223 is "FACE"
	// On a mobile keypad :)
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("PORT", "3223")

	m := martini.Classic()
	m.Get("/incoming", newCaller)
	m.Post("/incoming", newCaller)
	m.Get("/sections", readSections)
	m.Get("/threads/:handler", readThread) // Because threads shift around the :handler is needed.
	m.Get("/threads/:handler/:postnumber", readThreadPostNum)

	m.Run()
}
