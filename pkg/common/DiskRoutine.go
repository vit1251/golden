package commonfunc

import (
	"os"
	"os/user"
	"path"
)

func GetHomeDirectory() (string, error) {
	usr, err1 := user.Current()
	if err1 != nil {
		return "", err1
	}
	userHomeDir := usr.HomeDir
	return userHomeDir, nil
}

func GetLogDirectory() string {
	return GetFidoDirectory()
}

func checkDirectory(dir string) {
	if _, err1 := os.Stat(dir); err1 != nil {
		if os.IsNotExist(err1) {
			os.MkdirAll(dir, os.ModeDir|0755)
		}
	}
}

func checkDirectoryStructure(baseDirectory string) error {

	/* Check */
	checkDirectory(path.Join(baseDirectory, "Inbound"))
	checkDirectory(path.Join(baseDirectory, "Outbound"))

	checkDirectory(path.Join(baseDirectory, "TempInbound"))
	checkDirectory(path.Join(baseDirectory, "TempOutbound"))
	checkDirectory(path.Join(baseDirectory, "Temp"))

	checkDirectory(path.Join(baseDirectory, "FileBox"))

	return nil

}

func GetFidoDirectory() string {

	/* Check portable version execute */
	if currentDirectory, err1 := os.Getwd(); err1 == nil {
		newRoot := path.Join(currentDirectory, "Fido")
		if _ , err2 := os.Stat(newRoot); err2 == nil {
			checkDirectoryStructure(newRoot)
			return newRoot
		}
	}

	/* Check classic version execute */
	if homeDirectory, err1 := GetHomeDirectory(); err1 == nil {
		newRoot := path.Join(homeDirectory, "Fido")
//		if _ , err2 := os.Stat(newRoot); err2 == nil {
		checkDirectoryStructure(newRoot)
		return newRoot
//		}
	}

	return "Fido"
}

func GetInboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, "Inbound")
}

func GetFilesDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, "Files")
}

func GetOutboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, "Outbound")
}

func GetTempOutboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, "TempOutbound")
}

func GetTempInboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, "TempInbound")
}

func GetTempDirectory() string{
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, "Temp")
}
