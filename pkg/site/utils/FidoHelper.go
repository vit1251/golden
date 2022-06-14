package utils

import (
	"fmt"
	"github.com/vit1251/golden/pkg/packet"
)

func CreateNetAddr(source string) (string, error) {
	addr := packet.NewNetAddr()
	err1 := addr.SetAddr(source)
	if err1 != nil {
		return "", err1
	}
	result := fmt.Sprintf("f%d.n%d.z%d.binkp.net", addr.Node, addr.Net, addr.Zone)
	return result, nil
}
