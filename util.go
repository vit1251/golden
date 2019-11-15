package main

import (
	"bytes"
	"log"
)

func makeString(zstr []byte) string {
	var res string
	n := bytes.IndexByte(zstr, 0)
	log.Printf("zeroIndex = %d", n)
	if n > 0 {
		res = string(zstr[:n])
	} else {
		res = string(zstr[:])
	}
	return res
}