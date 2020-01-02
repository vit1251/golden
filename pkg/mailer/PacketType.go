package mailer

type PacketType uint8

const (
	CommandPacket  PacketType  = 0x00
	BinaryPacket   PacketType  = 0x01
)

type MessageType uint8

const (
	TextFrame     MessageType  = 0x00
	AddrFrame     MessageType  = 0x01
	PasswordFrame MessageType  = 0x02
	OkFrame       MessageType  = 0x03
	FileFrame     MessageType  = 0x04
	EndFrame      MessageType  = 0x05
	GotFrame      MessageType  = 0x06
	ErrorFrame    MessageType  = 0x07
	BusyFrame     MessageType  = 0x08
	GetFrame      MessageType  = 0x09
	SkipFrame     MessageType  = 0x0A
)