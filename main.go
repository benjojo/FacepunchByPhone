package main

import (
	"log"
	"os"
)

var info *log.Logger
var debug *log.Logger

func main() {
	info = log.New(os.Stdout, "[info] ", log.Ltime)
	info = log.New(os.Stdout, "[debug] ", log.Ltime|log.Lshortfile)
}
