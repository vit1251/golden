package widgets

import (
	"bytes"
	"io"
)

type ByteBuilder struct {
	io.Writer
	chunks [][]byte
	size   int
}

func NewByteBuilder() *ByteBuilder {
	newByteBuilder := new(ByteBuilder)
	//newByteBuilder.chunks = make([][]byte)
	return newByteBuilder
}

func (self *ByteBuilder) Append(chunk []byte) {
	//log.Printf("ByteBuilder: Append: chunk = %s", chunk)
	newChunk := make([]byte, len(chunk))
	size := copy(newChunk, chunk)
	self.chunks = append(self.chunks, newChunk)
	self.size += size
}

func (self ByteBuilder) Byte() []byte {
	return bytes.Join(self.chunks, []byte(""))
}

func (self *ByteBuilder) AppendString(str string) {
	self.Append([]byte(str))
}

func (self *ByteBuilder) Write(p []byte) (int, error) {
	self.Append(p)
	var size int = len(p)
	return size, nil
}
