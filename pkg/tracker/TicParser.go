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
	registry *registry.Container
}

func NewTicParser(r *registry.Container) *TicParser {
	tp := new(TicParser)
	tp.registry = r
	return tp
}

func (self *TicParser) prcessLine(ticFile TicFile, newLine string) (TicFile, error) {

	parts := strings.SplitN(newLine, " ", 2)
	partCount := len(parts)

	if partCount == 2 {

		var name string = parts[0]
		var value string = parts[1]

		/* Trim parameters */
		name = strings.Trim(name, " \t")
		value = strings.Trim(value, " \t")

		/* Process directive */
		if strings.EqualFold(name, "Area") {
			ticFile.SetArea(value)
		} else if strings.EqualFold(name, "Desc") {
			ticFile.SetDesc(value)
		} else if strings.EqualFold(name, "File") {
			ticFile.SetFile(value)
		} else if strings.EqualFold(name, "LFile") {
			ticFile.SetLFile(value)
		} else if strings.EqualFold(name, "Date") {
			if unixTime, err := strconv.ParseInt(value, 10, 64); err == nil {
				ticFile.SetUnixTime(unixTime)
			}
		} else {
			log.Printf("Unknown TIC directive: name = %+v value = %+v", name, value)
		}

	} else {
		log.Printf("Unknown TIC line: line = %+v", newLine)
	}

	return ticFile, nil
}

func (self *TicParser) Parse(stream io.Reader) (*TicFile, error) {

	charsetManager := self.restoreCharsetManager()

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
		*ticFile, _ = self.prcessLine(*ticFile, row)
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

func (self TicParser) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}