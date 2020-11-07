package mapper

type Stat struct {
	TicReceived      int
	TicSent          int
	EchomailReceived int
	EchomailSent     int
	NetmailReceived  int
	NetmailSent      int

	PacketReceived   int
	PacketSent       int

	MessageReceived  int
	MessageSent      int

	SessionIn        int
	SessionOut       int
}

