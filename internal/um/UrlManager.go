package um

import (
	"fmt"
	"github.com/vit1251/golden/pkg/registry"
	"net/url"
	"strings"
)

type UrlManager struct {
	registry.Service
}

type Url struct {
	pattern string
	params  map[string]string
}

func NewUrlManager(r *registry.Container) *UrlManager {
	new_um := new(UrlManager)
	new_um.SetRegistry(r)
	return new_um
}

func (self *UrlManager) CreateUrl(pattern string) *Url {
	new_url := new(Url)
	new_url.pattern = pattern
	new_url.params = make(map[string]string)
	return new_url
}

func (self *Url) SetParam(name string, value string) *Url {
	self.params[name] = value
	return self
}

func (self *Url) Build() string {
	var out string = self.pattern
	for k, v := range self.params {
		k = url.PathEscape(k)
		param_name := fmt.Sprintf("{%s}", k)
		out = strings.Replace(out, param_name, v, -1)
	}
	return out
}
