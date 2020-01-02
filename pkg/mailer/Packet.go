package mailer

type Packet struct {
	Type      PacketType    /* Packet type    */
	Payload []byte          /* Packet payload */
}
