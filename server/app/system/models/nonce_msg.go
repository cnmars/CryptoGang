package models

import(
	"math/rand"
	"time"
	"strconv"
)
type NonceMsg struct{
	Msg string `json:"msg"`
	Nonce string `json:"nonce"`
}
func DefaultNonceMsg() *NonceMsg {
	return &NonceMsg{
		Msg: "I confirm to sign with random nonce:",
	}
}
func (n *NonceMsg)GetNonce()string{
	return n.Nonce
}
func (n *NonceMsg)GetMsg()string{
	return n.Msg
}
func NewNonceMsg()NonceMsg{
	rand.Seed(time.Now().UnixNano())
	//生成六位随机数
	num := int64(rand.Intn(1000000))
	n :=NonceMsg{
		Msg:"I confirm to sign with random nonce:",
		Nonce:strconv.FormatInt(num , 10),
	}
	return n
}
