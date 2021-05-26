package chia_client

import (
	"encoding/json"
	"errors"
	"github.com/chuwt/fasthttp-client"
	"math/big"
)

type ChiaClient struct {
	url      string
	certPath string
	keyPath  string
}

func NewChiaClient(fullNodeUrl string, opts ...Opt) *ChiaClient {
	cc := new(ChiaClient)
	cc.url = fullNodeUrl
	for _, opt := range opts {
		opt(cc)
	}
	return cc
}

// GetCoins warn: maybe timeout if address has large spent and unspent tx from start to end
func (c *ChiaClient) GetCoins(req GetCoinsReq) (*GetCoinsRes, error) {
	if req.Start < 0 || req.Start > req.End {
		return nil, errors.New("start height must be greater than end height")
	} else if req.Address == "" {
		return nil, errors.New("address can't be empty")
	}
	puzzleHash, err := req.Address.PuzzleHash()
	if err != nil {
		return nil, err
	}

	data := struct {
		PuzzleHash        string `json:"puzzle_hash"`
		IncludeSpentCoins bool   `json:"include_spent_coins"`
		StartHeight       int64  `json:"start_height"`
		EndHeight         int64  `json:"end_height"`
	}{
		PuzzleHash:        puzzleHash.Hex(),
		IncludeSpentCoins: req.Spent,
		StartHeight:       req.Start,
		EndHeight:         req.End,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	res, err := fasthttp.NewClient().
		SetCrt(c.certPath, c.keyPath).
		AddBodyByte(dataBytes).
		AddHeader("content-type", "application/json").
		Post(c.url + "/get_coin_records_by_puzzle_hash")
	if err != nil {
		return nil, err
	}
	resp := new(GetCoinsRes)
	if err = json.Unmarshal(res.Body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type NewSignedTxReq struct {
	Coins      []*CoinRecord
	SendToList []*SendTo // you can send to multi address
	Fee        *big.Int
}

type SendTo struct {
	To     Address
	Amount *big.Int
}

// NewSignedTx generate a new signed tx
func (c *ChiaClient) NewSignedTx(req NewSignedTxReq) error {
	var maxSend, totalSend = big.NewInt(0), big.NewInt(0)
	// calculate max amount that address can send to
	for _, coin := range req.Coins {
		if !coin.Coin.Spent {
			maxSend = new(big.Int).Add(maxSend, coin.Coin.Amount)
		}
	}

	for _, send := range req.SendToList {
		totalSend = new(big.Int).Add(totalSend, send.Amount)
	}

	if maxSend.Cmp(totalSend) < 0 {
		return errors.New("insufficient balance")
	}

	// select coins from req.Coins that can spent for total send
	// sort first
	return nil

}

func (c *ChiaClient) newUnsignedTx() {

}

// SendTx send a tx to full_node by requesting /push_tx
func (c *ChiaClient) SendTx() {

}

type Opt func(*ChiaClient)

func TlsCertOpt(cerPath, keyPath string) Opt {
	return func(client *ChiaClient) {
		client.certPath = cerPath
		client.keyPath = keyPath
	}
}
