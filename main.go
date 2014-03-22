package main

import (
	"github.com/codegangsta/martini"
	"log"
	"os"
)

var info *log.Logger
var debug *log.Logger

func main() {
	info = log.New(os.Stdout, "[info] ", log.Ltime)
	debug = log.New(os.Stdout, "[debug] ", log.Ltime|log.Lshortfile)

	info.Println("Facepunch by phone, S16/03/2014")
	debug.Println("Debug text enabled")

	m := martini.Classic()
	m.Get("/incoming", func() string {
		return "Hello world!"
	})
	m.Run()
}
