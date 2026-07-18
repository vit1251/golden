package tracker

import (
	"bufio"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/registry"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type TicParserHandler func(string)

type TicParser struct {
    registry    *registry.Container
    handlers    map[string]func(*TicFile, string)
}

func NewTicParser(r *registry.Container) *TicParser {
    parser := &TicParser{
	registry: r,
	handlers: map[string]func(*TicFile, string){
	    "Area":   func(t *TicFile, v string) { t.SetArea(v) },
            "Desc":   func(t *TicFile, v string) { t.SetDesc(v) },
            "File":   func(t *TicFile, v string) { t.SetFile(v) },
	    "LFile":  func(t *TicFile, v string) { t.SetLFile(v) },
    	    "Date":   func(t *TicFile, v string) {
        	if unixTime, err := strconv.ParseInt(v, 10, 64); err == nil {
            	    t.SetUnixTime(unixTime)
        	}
    	    },
	    "Origin": func(t *TicFile, v string) { t.SetOrigin(v) },
	    "From":   func(t *TicFile, v string) { t.SetFrom(v) },
	    "To":     func(t *TicFile, v string) { t.SetTo(v) },
	    "Size":   func(t *TicFile, v string) {
        	if size, err := strconv.ParseInt(v, 10, 64); err == nil {
		    t.SetSize(size)
        	}
    	    },
    	    "Crc":    func(t *TicFile, v string) { t.SetCrc(v) },
    	    "LDesc":  func(t *TicFile, v string) { t.SetLDesc(v) },
	},
    }
    return parser
}

func (p *TicParser) prcessLine(ticFile *TicFile, newLine string) error {
    parts := strings.SplitN(newLine, " ", 2)
    if len(parts) != 2 {
	return nil
    }
    name := parts[0]
    value := parts[1]
    /* Trim parameters */
    name = strings.Trim(name, " \t")
    value = strings.Trim(value, " \t")
    /* Process directive */
    if handler, ok := p.handlers[name]; ok {
        handler(ticFile, value)
    } else {
	log.Printf("Unknown TIC directive: name = %+v value = %+v", name, value)
    }
    return nil
}

func (p *TicParser) Parse(stream io.Reader) (*TicFile, error) {
    charsetManager := charset.RestoreCharsetManager(p.registry)
    content, err1 := ioutil.ReadAll(stream)
    if err1 != nil {
	return nil, err1
    }
    newContent, err2 := charsetManager.DecodeMessageBody(content, "CP866")
    if err2 != nil {
	return nil, err2
    }
    rows := strings.Split(newContent, CRLF)
    ticFile := NewTicFile()
    for _, row := range rows {
	err3 := p.prcessLine(ticFile, row)
	if err3 != nil {
	    return ticFile, err3
	}
    }
    return ticFile, nil
}

func (self TicParser) ParseFile(filename string) (*TicFile, error) {

	/* Open FS object*/
	stream, err1 := os.Open(filename)
	if err1 != nil {
		return nil, err1
	}
	cacheStream := bufio.NewReader(stream)
	defer func() {
		stream.Close()
	}()

	/* Parse */
	ticFile, err2 := self.Parse(cacheStream)
	if err2 != nil {
		return nil, err2
	}

	return ticFile, nil
}
