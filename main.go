package main

import (
	"os"
	"log"
)

func main() {

	/* Parse command */
	var name string = "help"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	/* Process */
	if name == "pktinfo" {

		PktInfo()

	} else if name == "toss" {

		Tosser()

	} else if name == "mailer" {

		Mailer()

	} else if name == "reader" {

		Reader()

	} else {

		log.Printf("Usage: golden [command]")

	}

}
