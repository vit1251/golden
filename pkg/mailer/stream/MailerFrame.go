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
