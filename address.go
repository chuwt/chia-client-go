package chia_client

import (
	"encoding/hex"
	"errors"
)

type Address string

func (a Address) PuzzleHash() (PuzzleHash, error) {
	_, data, err := Decode(string(a))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("Invalid Address ")
	}
	decoded, err := ConvertBits(data, 5, 8, false)
	if err != nil {
		return nil, err
	}
	type Address string

	decodeBytes := make([]byte, len(decoded))
	for index, d := range decoded {
		decodeBytes[index] = uint8(d)
	}

	return decodeBytes, nil
}

type PuzzleHash []byte

func (ph PuzzleHash) Hex() string {
	return hex.EncodeToString(ph)
}
