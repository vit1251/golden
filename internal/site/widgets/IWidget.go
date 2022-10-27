package widgets

import (
	"io"
)

type IWidget interface {
	Render(w io.Writer) error
}
