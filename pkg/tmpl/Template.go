package tmpl

import (
	cmn "github.com/vit1251/golden/pkg/common"
	"runtime"
	"strings"
)

type Template struct {
	Vars map[string]string
}

func (self *Template) getPlatformName() string {
	var result string = runtime.GOOS
	result = strings.ToUpper(result[0:1]) + strings.ToLower(result[1:])
	return result
}

func (self *Template) getArchName() string {
	var result string = runtime.GOARCH
	return result
}

func NewTemplate() *Template {
	r := new(Template)
	r.Vars = make(map[string]string)

	platformName := r.getPlatformName()
	archName := r.getArchName()

	r.Set("GOLDEN_PLATFORM", platformName)
	r.Set("GOLDEN_ARCH", archName)
	r.Set("GOLDEN_VERSION", cmn.GetVersion())
	r.Set("GOLDEN_RELEASE_DATE", cmn.GetReleaseDate())
	r.Set("GOLDEN_RELEASE_HASH", "master")

	return r
}

func (self *Template) Set(name string, value string) {
	self.Vars[name] = value
}

func (self *Template) Render(msg string) (string, error) {

	newResult := msg

	for name, value := range self.Vars {
		varName := "{" + name + "}"
		newResult = strings.ReplaceAll(newResult, varName, value)
	}

	return newResult, nil
}