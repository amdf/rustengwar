package rustengwar

import "testing"

func TestConvert(t *testing.T) {
	russian := "русский текст"
	tengwar := Convert(russian)
	if russian != tengwar {
		t.Error("convert error")
	}
}
