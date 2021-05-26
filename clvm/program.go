package clvm

import (
	"encoding/hex"
	"io/ioutil"
)

const INFINITE_COST = 0x7FFFFFFFFFFFFFFF

var HiddenPuzzleHash, _ = hex.DecodeString("711d6c4e32c92e53179b199484cf8c897542bc57f2b22582799f9d657eec4699")

type Program struct {
	buf []byte
}

func LoadCLVM(path string) (*Program, error) {
	hexBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	buf, _ := hex.DecodeString(string(hexBytes))
	return &Program{
		buf: buf,
	}, nil
}
