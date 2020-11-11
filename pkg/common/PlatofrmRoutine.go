package commonfunc

import (
	"runtime"
	"strings"
)

func GetPlatform() string {
	var GoOS string = runtime.GOOS
	result := strings.ToUpper(GoOS[0:1]) + strings.ToLower(GoOS[1:])
	return result
}

func GetArch() string {
	var result string = runtime.GOARCH
	return result
}
