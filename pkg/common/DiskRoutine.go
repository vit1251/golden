package commonfunc

import (
	"os/user"
	"path"
)

func GetStorageDirectory() string {
	usr, err1 := user.Current()
	if err1 != nil {
		panic(err1)
	}
	userHomeDir := usr.HomeDir
	return userHomeDir
}

func GetLogDirectory() string {
	storageDirectory := GetStorageDirectory()
	debugDirectory := path.Join(storageDirectory, "Fido")
	return debugDirectory
}

