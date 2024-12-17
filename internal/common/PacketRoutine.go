package commonfunc

import (
	"fmt"
	"log"
	"time"
)

func MakePacketName() string {
	now := time.Now()
	unixTime := now.Unix()
	log.Printf("unixTime: dec = %d hex = %x", unixTime, unixTime)
	pktName := fmt.Sprintf("%08x.pkt", unixTime)
	log.Printf("pktName: name = %s", pktName)
	return pktName
}

func MakeTickName() string {
	now := time.Now()
	unixTime := now.Unix()
	log.Printf("unixTime: dec = %d hex = %x", unixTime, unixTime)
	pktName := fmt.Sprintf("%08x.tic", unixTime)
	log.Printf("pktName: name = %s", pktName)
	return pktName
}
