package file

import (
	"os"
	"bufio"
	"io"
	"log"
	"strings"
)

type TicParser struct {
	TicFile *TicFile
}

func NewTicParser() (*TicParser) {
	tp := new(TicParser)
	tp.TicFile = new(TicFile)
	return tp
}

func (self *TicParser) prcessLine(newLine string) {
	//
	parts := strings.SplitN(newLine, " ", 2)
	//
	log.Printf("parts = %+v", parts)
	//
	var key string = parts[0]
	var value string = parts[1]
	//
	if key == "Area" {
		self.TicFile.Area = value
	} else if key == "Desc" {
		self.TicFile.Desc = value
	} else if key == "File" {
		self.TicFile.File = value
	} else if key == "From" {
		self.TicFile.From = value
	} else if key == "To" {
		self.TicFile.To = value
	}



}

func (self *TicParser) Parse(stream io.Reader) (error) {

	cacheStream := bufio.NewReader(stream)

	for {
		newLine, err2 := cacheStream.ReadString('\n')
		if err2 == nil {
		} else if err2 == io.EOF {
			break
		} else {
			return err2
		}
		log.Printf("Line %s", newLine)
		newLine = strings.Trim(newLine, "\r\n ")
		self.prcessLine(newLine)
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
