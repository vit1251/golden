package packet

type Attribute int16

const (
	AttrPrivate      Attribute = 1 << 0
	AttrCrash        Attribute = 1 << 1
	AttrRecd         Attribute = 1 << 2
	AttrSend         Attribute = 1 << 3
	AttrFileAttached Attribute = 1 << 4
	AttrInTransit    Attribute = 1 << 5
	AttrOrphan       Attribute = 1 << 6
	AttrKillSent     Attribute = 1 << 7
	AttrLocal        Attribute = 1 << 8
)
