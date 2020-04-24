package tmpl

import (
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
	r.Set("GOLDEN_VERSION", "1.2.12")
	r.Set("GOLDEN_RELEASE_DATE", "2020-04-24 04:20 MSK")
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