package config

type EchoAreaType int

const (
	EchoAreaTypeNetmail       EchoAreaType    = 0
	EchoAreaTypeEcho          EchoAreaType    = 3
	EchoAreaTypeLocal         EchoAreaType    = 4
	EchoAreaTypeDupe          EchoAreaType    = 2
	EchoAreaTypeBad           EchoAreaType    = 1
	EchoAreaTypeNone          EchoAreaType    = 5
)
