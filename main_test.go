package rustengwar

import (
	"testing"
)

func TestInitDefault(t *testing.T) {
	var c Converter
	err := c.InitDefault()
	if err != nil {
		t.Error("default settings unmarshal error")
	}
}

func TestConvert(t *testing.T) {
	var c Converter
	var err error
	sampleString := "эльф"
	_, err = c.Convert(sampleString)
	if err == nil {
		t.Error("convert without initialization")
	}
	c.InitDefault()

	var tengwarStr string
	tengwarStr, err = c.Convert(sampleString)
	if err != nil {
		t.Error("convert failed")
	}

	if "`Vj´e" != tengwarStr {
		t.Error("failed: sample string")
	}
}
