package file

import (
	"bufio"
	"github.com/vit1251/golden/pkg/charset"
	"io"
	"log"
	"os"
	"strings"
)

type TicParserHandler func(string)

type TicParser struct {
	TicFile  *TicFile
	Handlers map[string]TicParserHandler
	cm       *charset.CharsetManager
}

func NewTicParser(cm *charset.CharsetManager) *TicParser {
	tp := new(TicParser)
	tp.cm = cm
	tp.TicFile = new(TicFile)
	tp.Handlers = make(map[string]TicParserHandler)
	tp.Handlers["Area"] = tp.processArea
	tp.Handlers["Desc"] = tp.processDesc
	tp.Handlers["File"] = tp.processFile
	tp.Handlers["From"] = nil
	tp.Handlers["To"] = nil
	tp.Handlers["Pw"] = nil
	tp.Handlers["File"] = tp.processFile
	tp.Handlers["Path"] = nil
	tp.Handlers["Crc"] = nil
	tp.Handlers["Size"] = nil
	tp.Handlers["Origin"] = nil
	tp.Handlers["LDesc"] = nil
	tp.Handlers["LDesc"] = nil
	tp.Handlers["Seenby"] = nil

	return tp
}

func (self *TicParser) prcessLine(newLine string) {
	//
	parts := strings.SplitN(newLine, " ", 2)
	//
	var key string = parts[0]
	var value string = parts[1]
	//
	var processCount int = 0
	for handerKey, handleService := range self.Handlers {
		if strings.EqualFold(handerKey, key) {
			if handleService != nil {
				handleService(value)
			}
			processCount += 1
		}
	}
	//
	if processCount == 0 {
		log.Printf("Unknown TIC keyword: %+v", parts)
	}

}

func (self *TicParser) Parse(stream io.Reader) error {

	cacheStream := bufio.NewReader(stream)
	for {
		newLine, err2 := cacheStream.ReadBytes('\n')
		if err2 == nil {
		} else if err2 == io.EOF {
			break
		} else {
			return err2
		}
		newRow, err3 := self.cm.Decode(newLine)
		if err3 != nil {
			return err3
		}
		newStr := string(newRow)
		newStr = strings.Trim(newStr, "\r\n ")
		self.prcessLine(newStr)
	}
	return nil
}

func (self *TicParser) ParseFile(filename string) (*TicFile, error) {

	/* Open FS object*/
	stream, err1 := os.Open(filename)
	if err1 != nil {
		return nil, err1
	}
	defer stream.Close()

	/* Parse */
	err2 := self.Parse(stream)
	if err2 != nil {
		return nil, err2
	}

	return self.TicFile, nil
}

func (self *TicParser) processArea(value string) {
	self.TicFile.Area = value
}

func (self *TicParser) processDesc(value string) {
	self.TicFile.Desc = value
}

func (self *TicParser) processFile(value string) {
	self.TicFile.File = value
}
