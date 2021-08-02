package rustengwar

import _ "embed" //embed?

//go:embed conv/default.json
var defaultConvJSON []byte

type defaultConversion struct {
	Replacements [][]string `json:"Replacements"`
}

//Convert russian string to tengwar string
func Convert(russian string) (tengwar string) {
	tengwar = russian //to be continued...
	//TODO: tengwar = strings.Replace(tengwar, "old", "new", -1)
	return
}
