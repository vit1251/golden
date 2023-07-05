package charset

import (
	"bytes"
	"golang.org/x/text/encoding/charmap"
	"testing"
)

func TestDecodeFromCP866(t *testing.T) {

	charsetManager := NewCharsetManager(nil)

	var msgBody []byte = []byte{0x8F, 0xE0, 0xA8, 0xA2, 0xA5, 0xE2, 0x2C, 0x20, 0xAC, 0xA8, 0xE0, 0x21}

	got, err1 := charsetManager.DecodeMessageBody(msgBody, "CP866")
	if err1 != nil {
		panic(err1)
	}

	var want string = "Привет, мир!"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestEncodeToCP850(t *testing.T) {

	charsetManager := NewCharsetManager(nil)

	var msgBody string = "Hallo Wereld!"

	got, err1 := charsetManager.EncodeMessageBody(msgBody, "CP850")
	if err1 != nil {
		panic(err1)
	}

	var want []byte = []byte{4861, 6c6c, 6f20, 5765, 7265, 6c64, 210a}
	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestEncodeToCP866(t *testing.T) {

	charsetManager := NewCharsetManager(nil)

	var msgBody string = "Привет, мир!"

	got, err1 := charsetManager.EncodeMessageBody(msgBody, "CP866")
	if err1 != nil {
		panic(err1)
	}

	var want []byte = []byte{0x8F, 0xE0, 0xA8, 0xA2, 0xA5, 0xE2, 0x2C, 0x20, 0xAC, 0xA8, 0xE0, 0x21}
	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestDecodeFromUTF8(t *testing.T) {

	charsetManager := NewCharsetManager(nil)

	var msgBody []byte = []byte{0xD0, 0x9F, 0xD1, 0x80, 0xD0, 0xB8, 0xD0, 0xB2, 0xD0, 0xB5, 0xD1, 0x82, 0x2C, 0x20,
		0xD0, 0xBC, 0xD0, 0xB8, 0xD1, 0x80, 0x21}

	got, err1 := charsetManager.DecodeMessageBody(msgBody, "UTF-8")
	if err1 != nil {
		panic(err1)
	}

	var want string = "Привет, мир!"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestEncodeToUTF8(t *testing.T) {

	charsetManager := NewCharsetManager(nil)

	var msgBody string = "Привет, мир!"

	got, err1 := charsetManager.EncodeMessageBody(msgBody, "UTF-8")
	if err1 != nil {
		panic(err1)
	}

	var want []byte = []byte{0xD0, 0x9F, 0xD1, 0x80, 0xD0, 0xB8, 0xD0, 0xB2, 0xD0, 0xB5, 0xD1, 0x82, 0x2C, 0x20,
		0xD0, 0xBC, 0xD0, 0xB8, 0xD1, 0x80, 0x21}
	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestSearchCharmapByIndex(t *testing.T) {
	charsetManager := NewCharsetManager(nil)
	got := charsetManager.searchCharmapByIndex("CP866")
	want := charmap.CodePage866
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
