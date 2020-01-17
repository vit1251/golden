package file

import (
	"bytes"
	"testing"
)

func TestTicParserParse1(t *testing.T) {

	sampleTic := bytes.NewBufferString("From 2:5023/24" +
		"To 2:5023/24.3752" +
		"Pw qazqaz" +
		"File AP200111.ZIP" +
		"Area NASA" +
		"Desc NASA Astronomy Picture of the Day (plus published report)" +
		"LDesc NASA Astronomy Picture of the Day" +
		"LDesc (plus published report)" +
		"Origin 1:153/757" +
		"Size 2770507" +
		"Crc 8A99C737" +
		"Path 1:153/757 1578729798 Sat Jan 11 08:03:18 2020 UTC htick/lnx 1.9.0-cur 2019-12-05" +
		"Path 1:261/38.0 @200111031003 EST+5" +
		"Path 2:5020/1042 1578730636 Sat Jan 11 08:17:16 2020 UTC htick/lnx 1.9.0-cur 20-04-17" +
		"Path 2:5034/10 1578731434 Sat Jan 11 08:30:34 2020 UTC htick/fbsd 1.9.0-cur 2019-01-08" +
		"Path 2:5023/24 1578762108 Sat Jan 11 12:01:48 2020 [Filin+/32 1.7b]")

	ticParser :=  NewTicParser()
	err := ticParser.Parse(sampleTic)
	if err != nil {
		t.Errorf("msg = %+v", err)
	}

}

