package packet

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var packetReaderNewPath string = "C:\\Users\\vit12\\Fido\\TempOutbound\\5fa748b6.pkt"

func TestPacketReaderNewReadPacketHeader(t *testing.T) {

	stream, err1 := os.Open(packetReaderNewPath)
	if err1 != nil {
		t.Errorf("i/o error on open: err = %+v", err1)
		return
	}
	defer stream.Close()

	cacheReader := bufio.NewReader(stream)

	packet, err2 := ioutil.ReadAll(cacheReader)
	if err2 != nil {
		t.Errorf("i/o error on read packet: err = %+v", err2)
		return
	}

	memReader := bytes.NewReader(packet)

	pktReader2 := NewPacketReaderNew(memReader)
	pktHeader , err3 := pktReader2.ReadPacketHeader()
	if err3 != nil {
		t.Errorf("packet read error on read packet header: err = %+v", err3)
		return
	}

	t.Logf("pktHeader = %+v", pktHeader)

	msgHeader, err4 := pktReader2.ReadMessageHeader()
	if err4 != nil {
		t.Errorf("packet read error on read message header: err = %+v", err4)
		return
	}

	t.Logf("msgHeader = %+v", msgHeader)

}

func BenchmarkPacketReaderNewReadPacketHeader(b *testing.B) {

	stream, err1 := os.Open(packetReaderNewPath)
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

		pktReader2 := NewPacketReaderNew(memReader)
		/* pktHeader , _ := */ pktReader2.ReadPacketHeader()

		//b.Logf("pktHeader = %+v", pktHeader)
	}

}
