package tmpl

import (
	"github.com/vit1251/golden/internal/common"
	"strings"
)

type Template struct {
	Vars map[string]string
}

func NewTemplate() *Template {
	r := new(Template)
	r.Vars = make(map[string]string)

	r.Set("GOLDEN_PLATFORM", commonfunc.GetPlatform())
	r.Set("GOLDEN_ARCH", commonfunc.GetArch())
	r.Set("GOLDEN_VERSION", commonfunc.GetVersion())
	r.Set("GOLDEN_RELEASE_DATE", commonfunc.GetReleaseDate())
	r.Set("GOLDEN_RELEASE_HASH", commonfunc.GetReleaseBranch())

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
