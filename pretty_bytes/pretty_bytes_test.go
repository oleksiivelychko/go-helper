package pretty_bytes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestPrintPrettyBytesToJSON(t *testing.T) {
	b := []byte(`{"hello":"world"}`)
	prettyJson := PrettyBytes(b, "	")

	fmt.Printf("%s\n", prettyJson)

	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, prettyJson); err != nil {
		t.Error(err)
	}

	if buffer.String() != string(b) {
		t.Errorf("[PrettyBytes(b []byte, indent string) []byte] -> `%s` != `%s`", buffer.String(), b)
	}
}
