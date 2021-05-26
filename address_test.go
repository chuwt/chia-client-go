package chia_client

import (
	"testing"
)

func TestAddress(t *testing.T) {
	address := Address("xch1f0ryxk6qn096hefcwrdwpuph2hm24w69jnzezhkfswk0z2jar7aq5zzpfj")
	puzzleHash, err := address.PuzzleHash()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("hex", puzzleHash.Hex())
}
