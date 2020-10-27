package commonfunc

import (
	"log"
	"os/user"
	"runtime"
	"strconv"
	"time"
)

func GetPlatform() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	} else if runtime.GOOS == "linux" {
		return "Linux"
	}
	return "Unknown"
}

func GetTime() string {
	// Sun, 26 Jan 2020 18:02:17 +0300
	now := time.Now().Format(time.RFC1123Z)
	log.Printf("Time is %s", now)
	return now
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
