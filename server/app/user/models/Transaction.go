package models

import (
	"math/big"
	// "github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	From     string
	To       string
	Value    *big.Int
	Gas      uint64
	GasPrice uint64
	Nonce    uint64
	Data     []byte
	Status   uint64//1 成功，0 失败
	Hash     string
	IsPending bool
}
