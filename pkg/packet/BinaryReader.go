package packet

import (
	"io"
//	"log"
	"encoding/binary"
)

type BinaryReader struct {
	reader io.Reader
	offset int64
}

func NewBinaryReader(reader io.Reader) *BinaryReader {
	binaryReader := new(BinaryReader)
	binaryReader.reader = reader
	binaryReader.offset = 0
	return binaryReader
}

func (self *BinaryReader) Offset() int64 {
	return self.offset
}

func (self *BinaryReader) ReadUINT8() (uint8, error) {
	var i uint8
	err := binary.Read(self.reader, binary.LittleEndian, &i)
	self.offset += 1
	return i, err
}

func (self *BinaryReader) ReadUINT16() (uint16, error) {
	var i uint16
	err := binary.Read(self.reader, binary.LittleEndian, &i)
	self.offset += 2
	return i, err
}

func (self *BinaryReader) ReadBytes(size int) ([]byte, error) {
	var i byte
	var cache []byte
	for j := 0; j < size; j++ {
		err := binary.Read(self.reader, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		cache = append(cache, i)
		self.offset += 1
	}
	return cache, nil
}

func (self *BinaryReader) ReadZString() ([]byte, error) {
	return self.ReadUntil('\x00')
}

func (self *BinaryReader) ReadUntil(ch byte) ([]byte, error) {
	var i byte
	var cache []byte
	for {
		err := binary.Read(self.reader, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		self.offset += 1
		if i == ch {
			break
		}
		cache = append(cache, i)
	}
	return cache, nil
}
