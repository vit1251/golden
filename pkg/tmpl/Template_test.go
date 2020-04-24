package tmpl

import (
	"log"
	"testing"
)

func TestMessageReplyTransformer1(t *testing.T) {

	tmpl := NewTemplate()

	str := "Golden/{GOLDEN_PLATFORM}-{GOLDEN_ARCH} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})"
	newVal, _ := tmpl.Render(str)
	log.Printf("value = %+v", newVal)

}