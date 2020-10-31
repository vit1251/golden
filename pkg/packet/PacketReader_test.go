package packet

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkToPackerReadPacketHeader(b *testing.B) {

	packetPath := "C:\\Users\\vit12\\Fido\\TempInbound\\9c6b8409.pkt"

	stream, _ := os.Open(packetPath)
	defer stream.Close()

	cacheReader := bufio.NewReader(stream)

	packet, _ := ioutil.ReadAll(cacheReader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		b.StopTimer()

		memReader := bytes.NewReader(packet)

		b.StartTimer()

		pktReader := NewPacketReader(memReader)
		/* pktHeader , _ := */ pktReader.ReadPacketHeader()

		//b.Logf("pktHeader = %+v", pktHeader)
	}

}
