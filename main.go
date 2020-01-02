package main

import (
	"os"
	"log"
)

func main() {

	if os.Args[1] == "pktinfo" {

		PktInfo()

	} else if os.Args[1] == "toss" {

		Tosser()

	} else if os.Args[1] == "reader" {

		Reader()

	} else {

		log.Printf("Usage: golden [command]")

	}

}
