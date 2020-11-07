package tracker

import (
	"bytes"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/registry"
	"testing"
)

func TestTicParserParse1(t *testing.T) {

	/* Create new registry */
	r := registry.NewContainer()
	r.Register("CharsetManager", charset.NewCharsetManager(r))

	/* Create TIC description */
	sampleTic := bytes.NewBuffer(nil)
	sampleTic.WriteString("From 2:5023/24" + CRLF)
	sampleTic.WriteString("To 2:5023/24.3752" + CRLF)
	sampleTic.WriteString("Pw 1111" + CRLF)
	sampleTic.WriteString("File AP200111.ZIP" + CRLF)
	sampleTic.WriteString("Area NASA" + CRLF)
	sampleTic.WriteString("Desc NASA Astronomy Picture of the Day (plus published report)" + CRLF)
	sampleTic.WriteString("LDesc NASA Astronomy Picture of the Day" + CRLF)
	sampleTic.WriteString("LDesc (plus published report)" + CRLF)
	sampleTic.WriteString("Origin 1:153/757" + CRLF)
	sampleTic.WriteString("Size 2770507" + CRLF)
	sampleTic.WriteString("Crc 8A99C737" + CRLF)
	sampleTic.WriteString("Path 1:153/757 1578729798 Sat Jan 11 08:03:18 2020 UTC htick/lnx 1.9.0-cur 2019-12-05" + CRLF)
	sampleTic.WriteString("Path 1:261/38.0 @200111031003 EST+5" + CRLF)
	sampleTic.WriteString("Path 2:5020/1042 1578730636 Sat Jan 11 08:17:16 2020 UTC htick/lnx 1.9.0-cur 20-04-17" + CRLF)
	sampleTic.WriteString("Path 2:5034/10 1578731434 Sat Jan 11 08:30:34 2020 UTC htick/fbsd 1.9.0-cur 2019-01-08" + CRLF)
	sampleTic.WriteString("Path 2:5023/24 1578762108 Sat Jan 11 12:01:48 2020 [Filin+/32 1.7b]" + CRLF)

	/* Parse */
	ticParser := NewTicParser(r)
	ticFile, err := ticParser.Parse(sampleTic)
	if err != nil {
		t.Errorf("msg = %+v", err)
	}

	t.Logf("t = %+v", ticFile)

}

