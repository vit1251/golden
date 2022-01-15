package commonfunc

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
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

const (
	InboundName      string = "Inbound"
	OutboundName     string = "Outbound"
	TempInboundName  string = "TempInbound"
	TempOutboundName string = "TempOutbound"
	TempName         string = "Temp"
	FileBoxName      string = "Files"
)

func checkDirectoryStructure(baseDirectory string) {

	/* Check */
	checkDirectory(path.Join(baseDirectory, InboundName))
	checkDirectory(path.Join(baseDirectory, OutboundName))

	checkDirectory(path.Join(baseDirectory, TempInboundName))
	checkDirectory(path.Join(baseDirectory, TempOutboundName))
	checkDirectory(path.Join(baseDirectory, TempName))

	checkDirectory(path.Join(baseDirectory, FileBoxName))

}

const (
	FidoName string = "Fido"
)

func GetFidoDirectory() string {

	/* Check portable version execute */
	if currentDirectory, err1 := os.Getwd(); err1 == nil {
		newRoot := path.Join(currentDirectory, FidoName)
		if _, err2 := os.Stat(newRoot); err2 == nil {
			checkDirectoryStructure(newRoot)
			return newRoot
		}
	}

	/* Check classic version execute */
	if homeDirectory, err1 := GetHomeDirectory(); err1 == nil {
		newRoot := path.Join(homeDirectory, FidoName)
		//		if _ , err2 := os.Stat(newRoot); err2 == nil {
		checkDirectoryStructure(newRoot)
		return newRoot
		//		}
	}

	return "Fido"
}

func GetInboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, InboundName)
}

func GetFilesDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, FileBoxName)
}

func GetOutboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, OutboundName)
}

func GetTempOutboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, TempOutboundName)
}

func GetTempInboundDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, TempInboundName)
}

func GetTempDirectory() string {
	fidoDirectory := GetFidoDirectory()
	return path.Join(fidoDirectory, TempName)
}

func GetPrevStorageFile() string {
	fidoDirectory := GetFidoDirectory()
	return filepath.Join(fidoDirectory, "..", "golden.sqlite3")
}

func GetModernStorageFile() string {
	fidoDirectory := GetFidoDirectory()
	return filepath.Join(fidoDirectory, "golden.sqlite3")
}
