package msg

import (
    "strings"
    "bytes"

    "github.com/vit1251/golden/pkg/charset"
)

func splitLines(data []byte) [][]byte {
    var lines [][]byte
    start := 0
    for i := 0; i < len(data); i++ {
        if data[i] == '\r' {
            lines = append(lines, data[start:i])
            if i+1 < len(data) && data[i+1] == '\n' { i++ }
            start = i + 1
        } else if data[i] == '\n' {
            lines = append(lines, data[start:i])
            start = i + 1
        }
    }
    if start < len(data) {
        lines = append(lines, data[start:])
    }
    return lines
}

type Kludge struct {
    Key   string
    Value string
}

// parseKludge разбирает строку вида \x01KEY: value или \x01Via ...
func parseKludge(line []byte) Kludge {

    // Проверяем длину входных данных
//    if len(line) == 0 {
//        return 
//    }

    // Отрезаем \x01
    line = line[1:]

    var key, val string

    // Ищем первый ':' или пробел
    idx := -1
    for i, b := range line {
        if b == ':' || b == ' ' {
            idx = i
            break
        }
    }

    if idx == -1 {
        key = string(line)
        val = ""
    } else {
        key = string(line[:idx])
        val = strings.TrimSpace(string(line[idx+1:]))
    }

    return Kludge{Key: key, Value: val}
}


// parseAreaName проверяет, является ли строка AREA-заголовком
func parseAreaName(line []byte) (string, bool) {
    if len(line) < 5 {
        return "", false
    }
    // Сравниваем первые 4 байта с "AREA" (без учета регистра — все AREA пишутся заглавными)
    if !bytes.Equal(line[:4], []byte("AREA")) {
        return "", false
    }
    if len(line) > 4 && (line[4] == ':' || line[4] == ' ') {
        val := strings.TrimSpace(string(line[5:]))
        return val, true
    }
    return "", false
}

func Unmarshal(data []byte) (*Message, error) {
    lines := splitLines(data)

    var areaName string = ""
    var kludges []Kludge
    var bodyLines [][]byte

    for i, line := range lines {

        // Первая строка — AREA (без \x01)
        if i == 0 && len(line) > 0 {
            if name, ok := parseAreaName(line); ok {
                areaName = name
                continue
            }
        }

        // Кладжа
        if len(line) > 0 && line[0] == 0x01 {
            kl := parseKludge(line)
            kludges = append(kludges, kl)
            continue
        }

        // Строка завершения
        if bytes.HasPrefix(line, []byte(" * Origin")) {
            bodyLines = append(bodyLines, line)
            break
        }

        // Тело сообщения
        bodyLines = append(bodyLines, line)
    }

    // Определяем кодировку
    charsetName := "CP866" // по умолчанию используем CP866 кодировку
    for _, kl := range kludges {
        if kl.Key == "CHRS" {
            // "CP866 2" -> "CP866"
            parts := strings.Fields(kl.Value)
            if len(parts) > 0 {
                charsetName = strings.ToUpper(parts[0])
            }
            break
        }
    }

    // Собираем тело сообщения
    bodyRaw := bytes.Join(bodyLines, []byte{'\n'})

    // Декодируем сообщение
    cm := charset.NewCharsetManager(nil)
    body, err := cm.DecodeMessageBody(bodyRaw, charsetName)
    if err != nil {
        body = string(bodyRaw) // fallback — сырые байты
    }

    // Заполняем Message
    msg := NewMessage()
    msg.SetArea(areaName)
    for _, kl := range kludges {
        switch kl.Key {
        case "FROM":
            msg.SetFrom(kl.Value)
        case "TO":
            msg.SetTo(kl.Value)
        case "SUBJ":
            msg.SetSubject(kl.Value)
        case "MSGID":
            msg.SetMsgID(kl.Value)
        case "REPLY":
            msg.SetReply(kl.Value)
        case "REPLYADDR", "@REPLYADDR":
            msg.SetFromAddr(kl.Value)
        }
    }
    msg.SetContent(body)
    return msg, nil

}
