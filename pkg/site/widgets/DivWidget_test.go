package widgets

import "testing"

func TestDivWidget_Render(t *testing.T) {

	builder := NewByteBuilder()

	div := NewDivWidget()
	div.SetClass("my-class")
	div.SetContent("Hello, world!")

	err1 := div.Render(builder)
	if err1 != nil {
		t.Errorf("Fail in Render on div: err = %+v", err1)
	}

	content := builder.Byte()

	t.Logf("content = %s", content)

}
