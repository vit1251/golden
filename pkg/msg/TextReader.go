package msg

type TextReader struct {
}

func NewTextReader() *TextReader {
	tr := new(TextReader)
	return tr
}

func (self *TextReader) Process(content string, cb func(oneLine string)) {

	var oneLine string
	var newLine bool = false

	for _, ch := range content {

		if ch == '\x0D' {
			newLine = true
			continue
		}
		if ch == '\x0A' {
			newLine = true
			continue
		}
		if newLine {
			cb(oneLine)
			oneLine = ""
			newLine = false
		}

		oneLine = oneLine + string(ch)
	}
	cb(oneLine)
}
