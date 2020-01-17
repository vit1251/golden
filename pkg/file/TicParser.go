package file

import (
	"os"
	"bufio"
	"io"
	"log"
)

type TicParser struct {
}

func NewTicParser() (*TicParser) {
	tp := new(TicParser)
	return tp
}

func (self *TicParser) prcessLine(newLine string) {
	log.Printf("processLine(%s)", newLine)
	//strings.HasPrefix("")
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
		self.prcessLine(newLine)
	}

	return nil

}

func (self *TicParser) ParseFile(filename string) (error) {

	/* Open FS object*/
	stream, err1 := os.Open(filename)
	if err1 != nil {
		return err1
	}
	defer stream.Close()

	/* Parse */
	err2 := self.Parse(stream)
	if err2 != nil {
		return err2
	}

	return nil
}
