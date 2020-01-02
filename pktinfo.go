package main

import (
	"github.com/vit1251/golden/pkg/packet"
	"os"
	"log"
)

func PktInfo() {

	name := os.Args[1]

	packet, err1 := packet.NewPacketReader(name)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer packet.Close()

	/* Process message */
	var msgCount int = 0

	for msg := range packet.Scan() {

		log.Printf("Area: %s", msg.Area)
		log.Printf("From: %s", msg.From)
		log.Printf("To: %s", msg.To)
		log.Printf("Subject: %s", msg.Subject)
		log.Printf("Time: %q", msg.Time)
		log.Printf("Body: %s", msg.Content)

		/* Add message count */
		msgCount += 1

	}

	/* Show summary */
	log.Printf("Packet contain %d message(s)", msgCount)

}