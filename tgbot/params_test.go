package tgbot

import (
	"strings"
	"testing"
)

func TestParamsURLValuesOk(t *testing.T) {
	str := "abc def"
	params := Params{"someParam%#$": &str}
	values, err := params.URLValues()
	if err != nil {
		t.Fatal("params.URLValues() failed: " + err.Error())
	}

	encodedParams := values.Encode()
	if strings.Compare(encodedParams, "someParam%25%23%24=abc+def") != 0 {
		t.Fatal("Unexpected Representation of Url-Encoded Values: \"" + encodedParams + "\"")
	}
}
