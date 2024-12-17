package tmpl

import (
	"testing"
        "strings"
)

func TestMessageReplyTransformer1(t *testing.T) {

	tmpl := NewTemplate()

	str := "Golden/{GOLDEN_PLATFORM}-{GOLDEN_ARCH} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})"
	newVal, err := tmpl.Render(str)
	if err != nil {
	    t.Fatalf("expected=* got=%q error=%q", newVal, err)
	}

	if !strings.Contains(newVal, "Golden") {
		t.Fatalf("expected=Golden* got=%q", newVal)
	}

	if strings.Contains(newVal, "{GOLDEN_PLATFORM}") {
	    t.Fatalf("expected=!{GOLDEN_PLATFORM} got=%q", newVal)
	}

	if strings.Contains(newVal, "{GOLDEN_ARCH}") {
	    t.Fatalf("expected=!{GOLDEN_ARCH} got=%q", newVal)
	}

	if strings.Contains(newVal, "{GOLDEN_VERSION}") {
	    t.Fatalf("expected=!{GOLDEN_VERSION} got=%q", newVal)
	}

	if strings.Contains(newVal, "{GOLDEN_RELEASE_DATE}") {
	    t.Fatalf("expected=!{GOLDEN_RELEASE_DATE} got=%q", newVal)
	}

	if strings.Contains(newVal, "{GOLDEN_RELEASE_HASH}") {
	    t.Fatalf("expected=!{GOLDEN_RELEASE_HASH} got=%q", newVal)
	}

}
