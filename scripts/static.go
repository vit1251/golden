package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func checkAllowExt(ext string) bool {
	var result = false
	var suffixs []string = []string{".svg", ".css", ".js", ".png", ".jpg"}
	for _, suffix := range suffixs {
		if strings.HasSuffix(ext, suffix) {
			result = true
		}
	}
	return result
}

func main() {

	baseDir := "static"
	fs, _ := ioutil.ReadDir(baseDir)
	out, _ := os.Create("./pkg/assets/000_image.go")

	out.Write([]byte("package assets\n"))
	out.Write([]byte("\n"))
	out.Write([]byte("func init() {\n"))

	for _, f := range fs {
		if checkAllowExt(f.Name()) {
			out.Write([]byte("\tMain[\"" + f.Name() + "\"]" + " = []byte{ "))
			srcName := path.Join(baseDir, f.Name())
			f, _ := os.Open(srcName)
			stream := bufio.NewReader(f)
			for {
				b, err := stream.ReadByte()
				if err != nil {
					break
				}
				repr := fmt.Sprintf("0x%X, ", b)
				out.Write([]byte(repr))
			}
			out.Write([]byte("}\n"))
		}
	}

	out.Write([]byte("}\n"))

}
