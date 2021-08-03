package rustengwar

import (
	_ "embed" //embed?
	"encoding/json"
	"errors"
	"strings"
)

//go:embed conv/default.json
var defaultConvJSON []byte

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
	err = json.Unmarshal(defaultConvJSON, &c.replacements)
	return
}

//Convert russian string to tengwar string
func (c Converter) Convert(russian string) (tengwar string, err error) {
	if 0 == len(c.replacements) {
		err = errors.New("converter not ready")
		return
	}

	str := " " + russian //add initial space for proper conversion
	for _, x := range c.replacements {
		str = strings.Replace(str, x[0], x[1], -1)
	}

	tengwar = strings.TrimSpace(str)

	return
}
