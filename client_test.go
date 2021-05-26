package chia_client

import "testing"

func TestClient(t *testing.T) {
	client := NewChiaClient(
		"https://192.168.1.164:8555",
		TlsCertOpt(
			"/Volumes/hdd1000gb/workspace/src/chia-client/ssl/full_node/private_full_node.crt",
			"/Volumes/hdd1000gb/workspace/src/chia-client/ssl/full_node/private_full_node.key",
		),
	)
	coins, err := client.GetCoins(GetCoinsReq{
		Address: "xch1xklqzcm6uk8cyem2m8px65wac7g492p3rfq5q8y4g0dfrnkvngdsm35dv4",
		Start:   0,
		End:     8388607,
	})
	if err != nil {
		t.Log(err)
		return
	}
	for _, coin := range coins.CoinRecords {
		t.Log("amount:", coin.Coin.Amount.String(), "puzzle_hash:", coin.Coin.PuzzleHash)
	}
}
