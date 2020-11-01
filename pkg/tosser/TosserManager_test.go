package tosser

import "testing"

func TestAverage(t *testing.T) {
	manager := TosserManager{}
	srcOrigin := "Друзья не те, кого ты больше всех любишь, а те, кто первым приходит на помощь"
	newOrigin := manager.prepareOrigin(srcOrigin)
	newOriginLen := len([]rune(newOrigin))
	if newOriginLen > 79 {
		t.Errorf("Origin after prepare is %d", newOriginLen)
	}

}

func TestCRC32(t *testing.T) {
	crc32 := makeCRC32([]byte("The quick brown fox jumps over the lazy dog"))
	expect := "414FA339"
	if crc32 != expect {
		t.Logf("CRC32: our value = %q <- expect %q", crc32, expect)
		t.Errorf("Wrong CRC32 calculation")
	}
}