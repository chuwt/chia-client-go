package chia_client

import (
	"encoding/json"
	"math/big"
)

// GetCoinsReq
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
	Coin Coin `json:"coin"`
}

type SmallCoin struct {
	Amount         *big.Int `json:"amount"`
	ParentCoinInfo string   `json:"parent_coin_info"`
	PuzzleHash     string   `json:"puzzle_hash"`
}

type Coin struct {
	SmallCoin
	Coinbase            bool  `json:"coinbase"`
	ConfirmedBlockIndex int64 `json:"confirmed_block_index"`
	Spent               bool  `json:"spent"`
	SpentBlockIndex     int64 `json:"spent_block_index"`
	Timestamp           int64 `json:"timestamp"`
}

func (c *Coin) ToJson() string {
	jsonBytes, _ := json.Marshal(c)
	return string(jsonBytes)
}

// PushTxReq
type PushTxReq struct {
	SpendBundle SpendBundle `json:"spend_bundle"`
}

type SpendBundle struct {
	CoinSolutions       []CoinSolution `json:"coin_solutions"`
	AggregatedSignature string         `json:"aggregated_signature"`
}

type CoinSolution struct {
	Coin         SmallCoin `json:"coin"`
	PuzzleReveal string    `json:"puzzle_reveal"`
	Solution     string    `json:"solution"`
}

// SignTxReq
type SignTxReq struct {
	Sk         string      `json:"sk"`
	UnsignedTx SpendBundle `json:"spend_bundle"`
	MsgList    [][]byte    `json:"msg_list"`
	PkList     [][]byte    `json:"pk_list"`
}
