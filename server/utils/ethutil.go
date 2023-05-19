package utils

import (
	"context"
	"encoding/hex"
	"fmt"
	"golddigger/app/user/models"
	"log"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
)

// PublicKeyBytesToAddress ...
func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

// 检查地址是否是有效的以太坊地址
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// 检查地址是否为零地址
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// 将wei（整数）转换为小数。 第二个参数是小数位数。
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// 将小数转换为wei(整数）。 第二个参数是小数位数
// wei := util.ToWei(0.02, 18)
// fmt.Println(wei) // 20000000000000000
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// 根据燃气上限和燃气价格计算燃气花费。
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// 从签名中提取R，S和V值。
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

func WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}
func EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}
func GetTranscationByHash(client *ethclient.Client, hash string) (*models.Transaction, error) {
	// client, err := ethclient.Dial("https://cloudflare-eth.com")
	// if err != nil {
	//     log.Fatal(err)
	// }
	var er error
	var tran = new(models.Transaction)
	txHash := common.HexToHash(hash)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	er = err
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(isPending)

	if err != nil {
		log.Fatal(err)
		er = err
	} else {

		tran.Value = tx.Value()
		tran.Gas = tx.Gas()
		tran.GasPrice = tx.GasPrice().Uint64()
		tran.Nonce = tx.Nonce()
		tran.Data = tx.Data()
		tran.To = tx.To().Hex()
		tran.Data = tx.Data()
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			er = err
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.LatestSignerForChainID(chainID), tx.GasPrice()); err == nil {
			// if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), tx.GasPrice()); err == nil {
			tran.From = msg.From().Hex()
			fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			er = err
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
		tran.Status = receipt.Status
	}
	return tran, er
}
