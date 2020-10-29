package commonfunc

import (
	"os/user"
	"runtime"
	"strconv"
)

func GetPlatform() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	} else if runtime.GOOS == "linux" {
		return "Linux"
	}
	return "Unknown"
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GetVersion() string {
	return "1.2.15"
}

func ParseSize(value []byte) (int, error) {
	var str string = string(value)
	size, err1 := strconv.ParseUint(str, 10, 32)
	return int(size), err1
}

func GetReleaseDate() string {
	return "2020-10-25 18:48 MSK"
}

func GetStorageDirectory() string {
	usr, err1 := user.Current()
	if err1 != nil {
		panic(err1)
	}
	userHomeDir := usr.HomeDir
	return userHomeDir
}

func GetLogDirectory() string {
	return "."
}
