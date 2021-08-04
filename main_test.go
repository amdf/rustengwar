package rustengwar

import (
	"testing"
)

func TestInitDefault(t *testing.T) {
	var c Converter
	err := c.InitDefault()
	if err != nil {
		t.Error("default settings unmarshal")
	}
}

func TestInitFile(t *testing.T) {
	var c Converter
	err := c.Init("nonexistent file")
	if err == nil {
		t.Error("init from nonexistent file")
	}

	err = c.Init("./conv/default.yaml")
	if err != nil {
		t.Error("init from file")
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
		t.Error("convert")
	}

	if "`Vj´e" != tengwarStr {
		t.Error("sample string")
	}
}
