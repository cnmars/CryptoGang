package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
    "github.com/cosmos/go-bip39"
    "github.com/decred/dcrd/bech32"
    "github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	// "golang.org/x/crypto/sha3"
)

type Wallet struct {
	privateKey, publicKey, Address string
}

func (w Wallet) GetAddress() string {
	return w.Address
}
func GenerateCommonWallet() (Wallet, error) {
	var wallet Wallet

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return wallet, err
		// log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	wallet.privateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return wallet, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		// log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	wallet.publicKey = hexutil.Encode(publicKeyBytes)[4:]

	wallet.Address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
	return wallet, nil
}

func ATOMCreateHDWallet(graph string) (Wallet, error) {
	seed := bip39.NewSeed(graph, "")
	// seed := bip39.NewSeed("blast about old claw current first paste risk involve victory edit current", "")
	// fmt.Println("Seed: ", hex.EncodeToString(seed)) // Seed:  dd5ffa7088c0fa4c665085bca7096a61e42ba92e7243a8ad7fbc6975a4aeea1845c6b668ebacd024fd2ca215c6cd510be7a9815528016af3a5e6f47d1cca30dd
	var wallet Wallet
	master, ch := hd.ComputeMastersFromSeed(seed)
	path := "m/44'/118'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		log.Fatal(err)
		return wallet, err
	}
	// Derivation Path:  m/44'/118'/0'/0/0'
	// Private Key:  69668f2378b43009b16b5c6eb5e405d9224ca2a326a65a17919e567105fa4e5a
	wallet.privateKey = hex.EncodeToString(priv)

	var privKey = secp256k1.PrivKey(priv)
	pubKey := privKey.PubKey()
	wallet.publicKey = hex.EncodeToString(pubKey.Bytes())
	// Public Key:  03de79435cbc8a799efc24cdce7d3b180fb014d5f19949fb8d61de3f21b9f6c1f8

	decodeString, err := hex.DecodeString(fmt.Sprintf("04%x", pubKey.Bytes()))
	if err != nil {
		log.Fatal(err)
		return wallet, err
	}

	// Convert test data to base32:
	conv, err := bech32.ConvertBits(decodeString, 8, 5, true)
	if err != nil {
		log.Fatal(err)
		return wallet, err
	}
	encoded, err := bech32.Encode("atom", conv)
	if err != nil {
		log.Fatal(err)
		return wallet, err
	}
	wallet.Address = encoded
	return wallet, nil
	// Show the encoded data.
	// fmt.Println("Wallet Address:", encoded)
}
func ETHGenerateHDWallet(graph string) (Wallet, error) {
	// "tag volcano eight thank tide danger coast health above argue embrace heavy"
	var wallet Wallet
	mnemonic := graph
	_wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
		return wallet, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := _wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
		return wallet, err
	}
	wallet.Address = account.Address.Hex()
	pvk, err := _wallet.PrivateKeyBytes(account)
	if err != nil {
		return wallet, err
	}
	wallet.privateKey = hexutil.Encode(pvk)[2:]
	// fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
	puk, err := _wallet.PublicKeyBytes(account)
	if err != nil {
		return wallet, err
	}
	wallet.publicKey = hexutil.Encode(puk)[4:]

	return wallet, nil
	// path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	// account, err = _wallet.Derive(path, false)
	// if err != nil {
	// 	return wallet, err
	// }
	// fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}

func CreateKs(password string) {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	// password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func ImportKs() {
	file := "./tmp/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}





