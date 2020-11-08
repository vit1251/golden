package packet

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var packetReaderPath string = "C:\\Users\\vit12\\Fido\\TempOutbound\\5fa748b6.pkt"

func TestPacketReaderReadPacketHeader(t *testing.T) {

	stream, err1 := os.Open(packetReaderPath)
	if err1 != nil {
		t.Errorf("i/o error on open: err = %+v", err1)
		return
	}
	defer stream.Close()

	cacheReader := bufio.NewReader(stream)

	packet, _ := ioutil.ReadAll(cacheReader)

	memReader := bytes.NewReader(packet)

	pktReader := NewPacketReader(memReader)
	pktHeader , _ := pktReader.ReadPacketHeader()

	t.Logf("pktHeader = %+v", pktHeader)

	msgHeader, _ := pktReader.ReadMessageHeader()

	t.Logf("msgHeader = %+v", msgHeader)

}

func BenchmarkPacketReaderReadPacketHeader(b *testing.B) {

	stream, err1 := os.Open(packetReaderPath)
	if err1 != nil {
		b.Errorf("i/o error on open: err = %+v", err1)
		return
	}
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
