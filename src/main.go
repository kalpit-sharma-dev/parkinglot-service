package main

import (
	"log"
	"runtime/debug"

	"github.com/kalpit-sharma-dev/parkinglot-service/src/server"
)

func main() {

	defer func() {
		r := recover()
		if r != nil {
			log.Print(r, debug.Stack())
		}
	}()
	log.Print("INFO : starting the application ....")
	log.Print("INFO : parkinglot-service started ....")

	server.LoadRoute()

}
