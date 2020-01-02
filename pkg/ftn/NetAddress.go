package ftn

type NetAddress struct {
	Zone      string
	Net       string
	Node      string
	Point     string
}

func NewNetAddress() (*NetAddress) {
	na := new(NetAddress)
	return na
}
