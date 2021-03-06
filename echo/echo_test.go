package echo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFmtBytes(t *testing.T) {
	b := []byte(`{"hello":"world"}`)
	prettyJson := FmtBytes(b, "	")

	fmt.Printf("%s\n", prettyJson)

	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, prettyJson); err != nil {
		t.Error(err)
	}

	if buffer.String() != string(b) {
		t.Errorf("[FmtBytes(b []byte, indent string) []byte] -> `%s` != `%s`", buffer.String(), b)
	}
}
