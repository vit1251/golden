package packet

import (
	"fmt"
	"os"
	"testing"
)

func TestMessageBodyParser_Parse(t *testing.T) {

	var msg []byte = []byte("AREA:RU.GOLDEN\r" +
		"\x01TZUTC: 0300\r" +
		"Hello,\r" +
		"\r" +
		"I send you message\r" +
		"\r" +
		"begin 644 cat.txt\r" +
		"#0V%T\r" +
		"`\r" +
		"end\r" +
		"\r" +
		"This is CAT.TXT with similar content.\r" +
		"\r" +
		" * Origin: Cat Station (7:777/777.7)")

	mParser := NewMessageBodyParser()
	mBody, err1 := mParser.Parse(msg)
	if err1 != nil {
		t.Errorf("Fail in Parse on MessageBodyParser: err = %+v", err1)
	}
	t.Logf("msg = %+v", mBody)

	content := mBody.GetContent()
	t.Logf("content = %q", content)

	/* Check attachments */
	for aIndex, a := range mBody.GetAttachments() {
		t.Logf("---")
		t.Logf("attach #%d: size = %d", aIndex, a.Len())
		raw := a.GetData()
		t.Logf("body = %s", raw)
	}

}

func TestNewMessageBodyParser(t *testing.T) {

	var origMsg []byte = []byte("Hello,\r" +
		"\r" +
		"My new icon:\r" +
		"\r" +
		"begin 644 uue-file-format.png\r" +
		"MB5!.1PT*&@H````-24A$4@```!`````0\"`0```\"U^C?J`````G-\"250(\"%7L\r" +
		"M1@0````)<$A9<P```&<```!G`=4HU\"8````9=$58=%-O9G1W87)E`'=W=RYI\r" +
		"M;FMS8V%P92YO<F>;[CP:````V$E$050H4XW0H4I#`12`X2]8QL\"@Y2+&A15E\r" +
		"MP7\"CMH%%HT46?`$?0S`YU]<TSS*9P2B(,$0V46PB`QF*86'\"CF&3*>R*_'#@\r" +
		"M'+YTA.!0_5>/EL.X\\:A_KY-CSZWYOT$X-Y<-GH50S09+BHI.)D!!2WE&+85`\r" +
		"MWD#7NZ$;'SK\"DU=WKEP;R),(94T]B4OK0D5=:MN^D&2#/<=3L.O4O9*.DK\"E\r" +
		"M*77@8@K.K-IP)*S9L6(HM:GB4T+.F]#W(H21KI'PH*VM+R=8T%\":4<-BQJ/B\r" +
		">Q_6?H*8]HUH(7P;'M@YA!3F=`````$E%3D2N0F\"\"\r" +
		"`\r" +
		"end\r" +
		"\r" +
		"Thanks.\r")

	mParser := NewMessageBodyParser()
	mBody, err1 := mParser.Parse(origMsg)
	if err1 != nil {
		t.Errorf("Fail in Parse on MessageBodyParser: err = %+v", err1)
	}
	t.Logf("msg = %+v", mBody)

	/* Check attachments */
	for aIndex, a := range mBody.GetAttachments() {
		t.Logf("---")
		t.Logf("attach #%d: size = %d", aIndex, a.Len())
		raw := a.GetData()
		if stream, err := os.Create(fmt.Sprintf("test_%d.png", aIndex)); err == nil {
			newBlock := raw
			t.Logf("newBlock: size = %d", len(newBlock))
			size, _ := stream.Write(newBlock)
			t.Logf("write: size = %d", size)
			stream.Close()
		}
		t.Logf("body = %s", raw)
	}

}
