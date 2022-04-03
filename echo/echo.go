package echo

import (
	"bytes"
	"encoding/json"
)

func FmtBytes(b []byte, indent string) []byte {
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", indent)

	return out.Bytes()
}
