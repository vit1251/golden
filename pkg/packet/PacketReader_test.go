package packet

import (
	//	"bufio"
	//	"bytes"
	//	"io/ioutil"
	//	"os"
	"testing"
)

//var packetReaderPath string = "C:\\Users\\vit12\\Fido\\TempOutbound\\5fa748b6.pkt"

func TestPacketReaderReadPacketHeader(t *testing.T) {

	//	stream, err1 := os.Open(packetReaderPath)
	//	if err1 != nil {
	//		t.Errorf("i/o error on open: err = %+v", err1)
	//		return
	//	}
	//	defer stream.Close()

	//	cacheReader := bufio.NewReader(stream)

	//	packet, _ := ioutil.ReadAll(cacheReader)

	//	memReader := bytes.NewReader(packet)

	//	pktReader := NewPacketReader(memReader)
	//	pktHeader , _ := pktReader.ReadPacketHeader()

	//	t.Logf("pktHeader = %+v", pktHeader)

	//	msgHeader, _ := pktReader.ReadPackedMessage()

	//	t.Logf("msgHeader = %+v", msgHeader)

}
