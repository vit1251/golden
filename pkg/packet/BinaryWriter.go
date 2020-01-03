package packet

import (
	"io"
//	"log"
	"encoding/binary"
)

type BinaryWriter struct {
	writer io.Writer
	offset int64
}

func NewBinaryWriter(writer io.Writer) (*BinaryWriter, error) {
	bw := new(BinaryWriter)
	bw.writer = writer
	bw.offset = 0
	return bw, nil
}

func (self *BinaryWriter) Offset() int64 {
	return self.offset
}

func (self *BinaryWriter) WriteUINT8(value uint8) (error) {
	err := binary.Write(self.writer, binary.LittleEndian, &value)
	self.offset += 1
	return err
}

func (self *BinaryWriter) WriteUINT16(value uint16) (error) {
	err := binary.Write(self.writer, binary.LittleEndian, value)
	self.offset += 2
	return err
}

//func (self *BinaryReader) ReadString(size int) ([]byte, error) {
//	var i byte
//	var cache []byte
//	for j := 0; j < size; j++ {
//		err := binary.Read(self.reader, binary.LittleEndian, &i)
//		if err != nil {
//			return nil, err
//		}
//		cache = append(cache, i)
//		self.offset += 1
//	}
//	var result string = string(cache)
//	log.Printf("ReadString(%d) = %v = %v", size, cache, result)
//	return cache, nil
//}
//
//func (self *BinaryReader) ReadUntil(ch byte) ([]byte, error) {
//	var i byte
//	var cache []byte
//	for {
//		err := binary.Read(self.reader, binary.LittleEndian, &i)
//		if err != nil {
//			return nil, err
//		}
//		cache = append(cache, i)
//		self.offset += 1
//		if i == ch {
//			break
//		}
//	}
//	var result string = string(cache)
//	log.Printf("ReadUntil(%c) = %v = %v", ch, cache, result)
//	return cache, nil
//}

func (self *BinaryWriter) Close() {
}
