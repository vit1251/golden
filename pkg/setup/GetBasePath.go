package setup

import (
	"log"
	"os/user"
	"path/filepath"
)

func GetBasePath() (string) {
	usr, err1 := user.Current()
	if err1 != nil {
		panic( err1 )
	}
	userHomeDirectory := usr.HomeDir
	log.Printf("userHomeDirectory = %+v", userHomeDirectory)

	newResult := filepath.Join(userHomeDirectory, "golden.sqlite3")

	return newResult
}
