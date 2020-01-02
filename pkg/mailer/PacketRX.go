package mailer

import (
	"log"
	"io"
	"encoding/binary"
	"bytes"
	"fmt"
)

const (
	FrameCommandMask uint16 = 0x08000
	FrameSizeMask    uint16 = 0x07FFF
)

func (self *Mailer) processAuthorization(chData []byte) {
	a := NewAuthorizer()

	a.SetChallengeData(string(chData))

	a.SetSecret("********")

	responseDigest, err := a.CalculateDigest()
	if err != nil {
		panic(err)
	}

	var password string = fmt.Sprintf("%s-%s-%s", "CRAM", "MD5", responseDigest)
	self.writePassword(password)

}

func (self *Mailer) processMessage(m *Message) {
	log.Printf("m = %q", m)

	/* Process OPT parameters */
	if bytes.HasPrefix(m.Body, []byte("OPT ")) {
		/* Split options */
		options := bytes.Fields([]byte(m.Body[4:]))
		log.Printf("options = %q", options)
		for _, option := range options {
			if bytes.HasPrefix(option, []byte("CRAM-MD5-")) {
				authDigest := option[9:]
				self.processAuthorization(authDigest)
			}
		}
	}
}

func (self *Mailer) processPacket(p *Packet) {
	/* Process state change packet */
	log.Printf("p = %q", p)
	if p.Type == CommandPacket {
		var messageType uint8 = p.Payload[0]
		log.Printf("messageType = %x", messageType)
		m := new(Message)
		m.Type = messageType
		m.Body = p.Payload[1:]
		self.processMessage(m)
	} else if p.Type == BinaryPacket {
		// TODO - continue write in stream ...
	}
}

func (self *Mailer) parsePacketHeader(packetHeader uint16) (PacketType, uint16, error) {

	var packetSize uint16 = packetHeader & FrameSizeMask
	var packetType PacketType
	if packetHeader & FrameCommandMask == FrameCommandMask {
		packetType = CommandPacket
	} else {
		packetType = BinaryPacket
	}

	return packetType, packetSize, nil

}

func (self *Mailer) processIncomingPacket() (*Packet, error) {

	p := new(Packet)

	var packetHeader uint16
	err1 := binary.Read(self.reader, binary.BigEndian, &packetHeader)
	if err1 == io.EOF {
		log.Printf("Session close.")
		return nil, err1
	}
	log.Printf("packetHeader = %04X", packetHeader)

	if packetType, packetSize, err := self.parsePacketHeader(packetHeader); err != nil {
		return nil, err
	} else {

		/* Get packet bytes */
		log.Printf("Packet RX with %d byte(s)", packetSize)
		mem := make([]byte, packetSize)
		_, err := io.ReadFull(self.reader, mem)
		if err != nil {
			return nil, err
		}

		/* Create packet */
		p.Type = packetType
		p.Payload = mem

	}

	return p, nil
}
