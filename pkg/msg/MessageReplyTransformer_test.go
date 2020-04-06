package msg

import (
	"log"
	"testing"
)

func TestMessageReplyTransformer1(t *testing.T) {

	var content string
	content += "Hi, Vitaliy!\n"
	content += "3 апр 20 13:48, Vitaliy Geydeko -> Stas Mishchenkov:\n"
	content += " SM>> Для чего нужны 200 минут на мобильные на линии с модемом?\n"
	content += " VG> ну я например с этой же линии звоню голосом.\n"
	content += "Я уже и не помню, как стационарным телефоном пользоваться без модема. ;)\n"

	mrt := NewMessageReplyTransformer()
	mrt.SetAuthor("VS")
	newContent := mrt.Transform(content)
	log.Printf("new = %+v", newContent)

}
