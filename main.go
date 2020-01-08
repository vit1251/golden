package main

import (
	"os"
	"time"
	"log"
)

func Periodic() {

	for {

		log.Printf("Check new mail")

		Mailer()

		Tosser()

		time.Sleep( 10 * time.Minute )

	}

}

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

	} else if name == "service" {

		go Periodic()
		Reader()  /* TODO - make sceduler or service manager ... */

	} else if name == "reader" {

		Reader()

	} else {

		log.Printf("Usage: golden [toss|mailer|reader|service]")

	}

}
