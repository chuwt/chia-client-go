package clvm

import (
	"encoding/hex"
	"testing"
)

func TestCLVM(t *testing.T) {
	program, err := LoadCLVM("./calculate_synthetic_public_key.clvm.hex")
	if err != nil {
		return
	}
	t.Log(hex.EncodeToString(program.buf))
}
