package chia_client

import (
	"encoding/json"
	"errors"
	"github.com/chuwt/fasthttp-client"
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

// SendTx send a tx to full_node by requesting /push_tx
func (c *ChiaClient) PushTx(req SpendBundleReq) ([]byte, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := fasthttp.NewClient().
		SetCrt(c.certPath, c.keyPath).
		AddBodyByte(data).
		AddHeader("content-type", "application/json").
		Post(c.url + "/push_tx")
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

type Opt func(*ChiaClient)

func TlsCertOpt(cerPath, keyPath string) Opt {
	return func(client *ChiaClient) {
		client.certPath = cerPath
		client.keyPath = keyPath
	}
}
