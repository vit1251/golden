package commonfunc

import (
	"os/user"
	"path"
)

func GetHomeDirectory() string {
	usr, err1 := user.Current()
	if err1 != nil {
		panic(err1)
	}
	userHomeDir := usr.HomeDir
	return userHomeDir
}

func GetLogDirectory() string {
	return GetFidoDirectory()
}

func GetFidoDirectory() string {
	storageDirectory := GetHomeDirectory()
	debugDirectory := path.Join(storageDirectory, "Fido")
	return debugDirectory
}
