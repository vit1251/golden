package commonfunc

import "runtime"

func GetPlatform() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	} else if runtime.GOOS == "linux" {
		return "Linux"
	}
	return "Unknown"
}

