package rustengwar

import (
	"encoding/json"
	"testing"
)

func TestConvert(t *testing.T) {
	russian := "русский текст"
	tengwar := Convert(russian)
	if russian != tengwar {
		t.Error("convert error")
	}
}

func TestLoadReplacements(t *testing.T) {
	var v defaultConversion
	err := json.Unmarshal(defaultConvJSON, &v)
	if err != nil {
		t.Error("Replacements json Unmarshal error")
	}
}
