package main

import (
	"github.com/vit1251/golden/pkg/packet"
	"os"
	"fmt"
)

func PktInfo() {

	name := os.Args[2]

	packet, err1 := packet.NewPacketReader(name)
	if err1 != nil {
		panic(err1)
	}
	defer packet.Close()

	/* Process message */
	var msgCount int = 0

	for msg := range packet.Scan() {

		fmt.Printf("From: \"%s\" <%s>\n", msg.From, "no-reply@fidonet.org")
		fmt.Printf("To: \"%s\" <%s>\n", msg.To, "no-reply@fidonet.org")
		fmt.Printf("Subject: %s\n", msg.Subject)
		fmt.Printf("Newsgroups: %s\n", msg.Area)
		fmt.Printf("Date: %s\n", msg.Time)
		fmt.Printf("\n")
		fmt.Printf("%s", msg.Content)

		/* Add message count */
		msgCount += 1

	}

}