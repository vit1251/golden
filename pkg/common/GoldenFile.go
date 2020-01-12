package common

import (
	"os"
	"bufio"
	"io"
	"log"
)

/*
From 2:5023/24
To 2:5023/24.3752
Pw qazqaz
File AP200111.ZIP
Area NASA
Desc NASA Astronomy Picture of the Day (plus published report)
LDesc NASA Astronomy Picture of the Day
LDesc (plus published report)
Origin 1:153/757
Size 2770507
Crc 8A99C737
Path 1:153/757 1578729798 Sat Jan 11 08:03:18 2020 UTC htick/lnx 1.9.0-cur 2019-12-05
Path 1:261/38.0 @200111031003 EST+5
Path 2:5020/1042 1578730636 Sat Jan 11 08:17:16 2020 UTC htick/lnx 1.9.0-cur 20-04-17
Path 2:5034/10 1578731434 Sat Jan 11 08:30:34 2020 UTC htick/fbsd 1.9.0-cur 2019-01-08
Path 2:5023/24 1578762108 Sat Jan 11 12:01:48 2020 [Filin+/32 1.7b]
*/

type GoldenFile struct {
	From     string     /* From */
	To       string
	File     string
	Area     string
	Desc     string
	Origin   string
	Size     string
	CRC      string
	Path   []string
}

func NewGoldenFile() (*GoldenFile) {
	gf := new(GoldenFile)
	return gf
}

func (self *GoldenFile) prcessLine(newLine string) {
	log.Printf("processLine(%s)", newLine)
	//strings.HasPrefix("")
}

func (self *GoldenFile) ParseTic(filename string) (error) {

	stream, err1 := os.Open(filename)
	if err1 != nil {
		return err1
	}
	defer stream.Close()

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
