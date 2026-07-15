package utils

import (
    "fmt"
    "strings"
)

func HexDumpGrouped(data []byte) string {
    var out strings.Builder
    for i := 0; i < len(data); i += 16 {
        out.WriteString(fmt.Sprintf("%08x  ", i))
        for j := 0; j < 16; j++ {
            if i+j < len(data) {
                out.WriteString(fmt.Sprintf("%02x", data[i+j]))
            } else {
                out.WriteString("  ")
            }
            if j%4 == 3 && j != 15 {
                out.WriteString("  ")
            } else if j != 15 {
                out.WriteByte(' ')
            }
        }
        out.WriteString("  |")
        for j := 0; j < 16 && i+j < len(data); j++ {
            b := data[i+j]
            if b >= 32 && b <= 126 {
                out.WriteByte(b)
            } else {
                out.WriteByte('.')
            }
        }
        out.WriteByte('|')
        if i+16 < len(data) {
            out.WriteByte('\n')
        }
    }
    return out.String()
}
