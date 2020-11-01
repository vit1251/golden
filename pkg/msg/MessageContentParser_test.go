package msg

import "testing"

func TestMessageContentParser_Parse_Windows(t *testing.T) {

	/* Source message on Windows machine */
	source :=
		"Hello," + CR + LF +
		CR + LF +
		"Sunday November 01 2020 09:01, from Nil Alexandrov -> Vitold Sedyshev:" + CR + LF +
		CR + LF +
		" VS>> Есть понимание из-за чего это произошло?" + CR + LF +
		" NA> Письма идентичные." + CR + LF +
		CR + LF +
		"Я думаю, что Дмитрия пока беспокоить по этому поводу рано." + CR + LF +
		CR + LF

	mp := NewMessageContentParser()
	mc, err := mp.Parse(source)
	if err != nil {
		t.Errorf("Fail in Parse on MessageContentParser: err = %+v", err)
	}

	t.Logf("mc = %+v", mc)

}
