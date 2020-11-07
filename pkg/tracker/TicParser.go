package tracker

import (
	"bufio"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/registry"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type TicParserHandler func(string)

type TicParser struct {
	ticFile  *TicFile
	handlers map[string]TicParserHandler
	registry *registry.Container
}

func NewTicParser(r *registry.Container) *TicParser {
	tp := new(TicParser)
	tp.registry = r
	tp.handlers = make(map[string]TicParserHandler)
	tp.initializeHandler()
	return tp
}

func (self *TicParser) registerHandler(name string, handler TicParserHandler) {
	self.handlers[name] = handler
}

func (self *TicParser) initializeHandler() {
	self.registerHandler("Area", self.processArea)
	self.registerHandler("Desc", self.processDesc)
	self.registerHandler("File", self.processFile)
	self.registerHandler("From", nil)
	self.registerHandler("To", nil)
	self.registerHandler("Pw", nil)
	self.registerHandler("File", self.processFile)
	self.registerHandler("Path", nil)
	self.registerHandler("Crc", nil)
	self.registerHandler("Size", nil)
	self.registerHandler("Origin", nil)
	self.registerHandler("LDesc", nil)
	self.registerHandler("LDesc", nil)
	self.registerHandler("Seenby", nil)
}

func (self *TicParser) prcessLine(newLine string) {

	parts := strings.SplitN(newLine, " ", 2)
	partCount := len(parts)

	if partCount == 2 {
		var key string = parts[0]
		var value string = parts[1]
		//
		var processCount int = 0
		for handerKey, handleService := range self.handlers {
			if strings.EqualFold(handerKey, key) {
				if handleService != nil {
					handleService(value)
				}
				processCount += 1
			}
		}

		if processCount == 0 {
			log.Printf("Unknown TIC keyword: %+v", parts)
		}

	} else {
		log.Printf("Unknown TIC directive: %+v", newLine)
	}

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

	self.ticFile = new(TicFile)

	for _, row := range rows {
		self.prcessLine(row)
	}

	return self.ticFile, nil
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

func (self *TicParser) processArea(value string) {
	self.ticFile.SetArea(value)
}

func (self *TicParser) processDesc(value string) {
	self.ticFile.Desc = value
}

func (self *TicParser) processFile(value string) {
	self.ticFile.File = value
}

func (self TicParser) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}