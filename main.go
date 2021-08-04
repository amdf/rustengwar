package rustengwar

import (
	_ "embed" //embed
	"errors"
	"strings"

	"gopkg.in/yaml.v2"
)

//go:embed conv/default.yaml
var defaultConvData []byte

//Converter (Russian to Tengwar)
type Converter struct {
	replacements [][]string
}

//InitDefault conversion scheme
func (c *Converter) InitDefault() (err error) {
	if nil == c {
		err = errors.New("Converter == nil")
		return
	}
	err = yaml.Unmarshal(defaultConvData, &c.replacements)
	return
}

//Convert russian string to tengwar string
func (c Converter) Convert(russian string) (tengwar string, err error) {
	if 0 == len(c.replacements) {
		err = errors.New("converter not ready")
		return
	}

	//TODO: check russian string contains cyrillic only

	str := " " + russian //add initial space for proper conversion
	for _, x := range c.replacements {
		str = strings.Replace(str, x[0], x[1], -1)
	}

	tengwar = strings.TrimSpace(str)

	return
}
