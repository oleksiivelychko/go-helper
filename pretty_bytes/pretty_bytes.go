package pretty_bytes

import (
	"bytes"
	"encoding/json"
)

func PrettyBytes(b []byte, indent string) []byte {
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", indent)

	return out.Bytes()
}
