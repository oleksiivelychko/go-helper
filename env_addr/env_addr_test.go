package env_addr

import (
	"os"
	"testing"
)

func TestGetAddr(t *testing.T) {
	_ = os.Setenv("HOST", "host.local")
	_ = os.Setenv("PORT", "8080")

	addr := GetAddr()

	if addr != "host.local:8080" {
		t.Errorf("[func GetAddr() string] -> %s != %s", addr, "host.local:8080")
	}
}
