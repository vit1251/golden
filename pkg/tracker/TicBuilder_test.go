package tracker

import (
	"encoding/hex"
	"testing"
)

func TestTicBuilder_Build(t *testing.T) {

	newTicBuilder := NewTicBuilder()
	newTicBuilder.SetFile("hello.txt")
	newTicBuilder.SetDesc("Hello, world")
	newTicBuilder.SetSize(13)
	newTicBuilder.SetPw("1111")

	content := newTicBuilder.Build()
	t.Logf("content = %+v", hex.Dump([]byte(content)))

}
