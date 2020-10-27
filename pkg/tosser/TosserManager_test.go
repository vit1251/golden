package tosser

import "testing"

func TestAverage(t *testing.T) {
	manager := TosserManager{}
	srcOrigin := "Друзья не те, кого ты больше всех любишь, а те, кто первым приходит на помощь"
	newOrigin := manager.prepareOrigin(srcOrigin)
	newOriginLen := len(newOrigin)
	if newOriginLen > 79 {
		t.Errorf("Origin after prepare is %d", newOriginLen)
	}

}
