package main

import (
	"github.com/vit1251/golden/pkg/packet"
	"io"
	"os"
	"fmt"
	"log"
)

func PktInfo() {

	name := os.Args[2]

	pr, err1 := packet.NewPacketReader(name)
	if err1 != nil {
		panic(err1)
	}
	defer pr.Close()

	/* Read pakcet header */
	pktHeader, err2 := pr.ReadPacketHeader()
	if err2 != nil {
		panic(err2)
	}
	log.Printf("pktHeader = %+v", pktHeader)

	/* Process message */
	var msgCount int = 0
	for {

		/* Read message header */
		msgHeader, err5 := pr.ReadMessageHeader()
		if err5 == io.EOF {
			break
		}
		if err5 != nil {
			panic(err5)
		}
		log.Printf("msgHeader = %+v", msgHeader)

		/* Read message body */
		rawBody, err6 := pr.ReadMessage()
		if err6 != nil {
			panic(err6)
		}

		/* Process message */
		msgParser, err7 := packet.NewMessageBodyParser()
		if err7 != nil {
			panic(err7)
		}
		msgBody, err8 := msgParser.Parse(rawBody)
		if err8 != nil {
			panic(err8)
		}

		/* Determine area */
		var areaName string = msgBody.GetArea()

		fmt.Printf("From: \"%s\" <%s>\n", msgHeader.FromUserName, msgHeader.OrigAddr.String() )
		fmt.Printf("To: \"%s\" <%s>\n", msgHeader.ToUserName, msgHeader.DestAddr.String() )
		fmt.Printf("Subject: %s\n", msgHeader.Subject)
		fmt.Printf("Newsgroups: %s\n", areaName)
		fmt.Printf("Date: %s\n", msgHeader.Time)
		fmt.Printf("\n")
		fmt.Printf("%s", msgBody.Body)

		/* Add message count */
		msgCount += 1

	}

}
