package chia_client

import "math/big"

type GetCoinsReq struct {
	Address Address
	Start   int64
	End     int64
	Spent   bool
}

type GetCoinsRes struct {
	CoinRecords []*CoinRecord `json:"coin_records"`
}

type CoinRecord struct {
	Coin struct {
		Amount              *big.Int `json:"amount"`
		ParentCoinInfo      string   `json:"parent_coin_info"`
		PuzzleHash          string   `json:"puzzle_hash"`
		Coinbase            bool     `json:"coinbase"`
		ConfirmedBlockIndex int64    `json:"confirmed_block_index"`
		Spent               bool     `json:"spent"`
		SpentBlockIndex     int64    `json:"spent_block_index"`
		Timestamp           int64    `json:"timestamp"`
	} `json:"coin"`
}
