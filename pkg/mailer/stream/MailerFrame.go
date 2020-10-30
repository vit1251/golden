package stream

type CommandFrame struct {
	CommandID CommandID /* Command ID  */
	Body      []byte
}

type DataFrame struct {
	Body []byte
}

type Frame struct {
	Command bool
	DataFrame
	CommandFrame
}

func (self Frame) IsCommandFrame() bool {
	return self.Command
}

func (self Frame) IsDataFrame() bool {
	return self.Command == false
}
