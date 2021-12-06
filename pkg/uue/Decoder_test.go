package uue

import (
	"testing"
	"bytes"
	"log"
)

func Test_Decoder_Decode(t *testing.T) {

	rawData := []byte("#0V%T")
	var out bytes.Buffer

	uueDecoder := NewDecoder(&out)
	err1 := uueDecoder.Decode(rawData)
	log.Printf("err = %+v", err1)

	log.Printf("out = %s", out.String())

}
