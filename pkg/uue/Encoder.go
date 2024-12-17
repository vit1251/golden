package uue

import (
	"io"
)

type Encoder struct {
}

func NewEncoder(w io.Writer) *Encoder {
	return new(Encoder)
}
