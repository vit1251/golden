package msg

type MessageTextProcessor struct {
}

func NewMessageTextProcessor() *MessageTextProcessor {
    mr := new(MessageTextProcessor)
    return mr
}

func StringSplitLines(s string) []string {
    var lines []string
    start := 0
    for i := 0; i < len(s); i++ {
	if s[i] == '\r' {
	    lines = append(lines, s[start:i])
	    if i+1 < len(s) && s[i+1] == '\n' {
		i++
	    }
	    start = i + 1
	} else if s[i] == '\n' {
	    lines = append(lines, s[start:i])
	    start = i + 1
	}
    }
    if start < len(s) {
	lines = append(lines, s[start:])
    }
    return lines
}

func (self *MessageTextProcessor) Prepare(data string) (*MessageDocument) {
    var items []MessageDocumentElement
    lines := StringSplitLines(data)
    for _, line := range lines {
        // Шаг 1. Парсер строки
        mlp := NewMessageLineParser()
	ml := mlp.Parse(line)
	// Шаг 2. Перекладываем результат
	items = append(items, MessageDocumentElement{
	    Author: ml.Author,
	    QuoteLevel: ml.QuoteLevel,
	    Text: ml.Text,
	})
    }
    return &MessageDocument{
        items: items,
    }
}
