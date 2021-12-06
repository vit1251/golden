package uue

import (
	"io"
	"log"
	"bytes"
)

type Decoder struct {
	writer		io.Writer        /* Decode data consumer    */
}

func decodeByte(b byte) byte {
	r := byte(b) - 32
	r = r & 077
	return r
}

func NewDecoder(w io.Writer) *Decoder {
	r := &Decoder{
		writer: w,
	}
	return r
}

const (
	UUE_DECODER_STATE_SIZE    = 1
	UUE_DECODER_STATE_BLOCK   = 2
)

func (self *Decoder) Decode(row []byte) error {

	stream := bytes.NewBuffer(row)
	var state int = UUE_DECODER_STATE_SIZE
	var n byte

	for {
		/* Extract UUE section count */
		if state == UUE_DECODER_STATE_SIZE {

			var sizeByte byte
			sizeByte, err1 := stream.ReadByte()
			if err1 != nil {
				return err1
			}

			n = decodeByte(sizeByte)
			log.Printf("UUE: n = %d", n)

			state = UUE_DECODER_STATE_BLOCK

		} else if state == UUE_DECODER_STATE_BLOCK {

			log.Printf("n = %d", n)

			if n >= 1 {
				rawByte0, _ := stream.ReadByte()
				rawByte1, _ := stream.ReadByte()

				b1 := decodeByte(rawByte0) << 2 | decodeByte(rawByte1) >> 4
				self.writer.Write([]byte{b1})

				if n >= 2 {
					rawByte2, _ := stream.ReadByte()

					b2 := decodeByte(rawByte1) << 4 | decodeByte(rawByte2) >> 2
					self.writer.Write([]byte{b2})

					if n >= 3 {
						rawByte3, _ := stream.ReadByte()

						b3 := decodeByte(rawByte2) << 6 | decodeByte(rawByte3)
						self.writer.Write([]byte{b3})

						n = n - 1
					}
					n = n - 1
				}
				n = n - 1
			} else {

				log.Printf("Line decode is complete: n = %d", n)
				break

			}

		}
	}

	return nil

}
