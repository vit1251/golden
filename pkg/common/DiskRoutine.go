package commonfunc

import "os/user"

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

