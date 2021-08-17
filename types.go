package chia_client

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"
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
	ParentCoinInfo string   `json:"parent_coin_info"`
	PuzzleHash     string   `json:"puzzle_hash"`
	Amount         *big.Int `json:"amount"`
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

func (sb *SpendBundle) TxHash() (string, error) {
	tmpBytes := make([]byte, 4)
	amountBytes := make([]byte, 8)
	txHashList := make([]string, 0)
	txHashList = append(txHashList, hex.EncodeToString(big.NewInt(int64(len(sb.CoinSolutions))).FillBytes(tmpBytes)))
	for _, cs := range sb.CoinSolutions {
		txHashList = append(txHashList, strings.TrimPrefix(cs.Coin.ParentCoinInfo, "0x"))
		txHashList = append(txHashList, strings.TrimPrefix(cs.Coin.PuzzleHash, "0x"))
		txHashList = append(txHashList, hex.EncodeToString(cs.Coin.Amount.FillBytes(amountBytes)))
		txHashList = append(txHashList, strings.TrimPrefix(cs.PuzzleReveal, "0x"))
		txHashList = append(txHashList, strings.TrimPrefix(cs.Solution, "0x"))
	}
	txHashList = append(txHashList, strings.TrimPrefix(sb.AggregatedSignature, "0x"))
	m, err := hex.DecodeString(strings.Join(txHashList, ""))
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(m)
	return hex.EncodeToString(hash[:]), nil
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
