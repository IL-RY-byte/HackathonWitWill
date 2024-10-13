package main

import (
	"flag"
	"go_project/common" // Adjust this import path to the actual package path
	"log"
)

func main() {
	mode := flag.String("mode", "", "Start in either 'client' or 'server' mode")
	flag.Parse()

	if *mode == "server" {
		log.Println("Running server...")
		common.StartServer()
	} else if *mode == "client" {
		log.Println("Running client...")
		common.StartClient()
	} else {
		log.Fatal("Please specify the mode: -mode client or -mode server")
	}
}
