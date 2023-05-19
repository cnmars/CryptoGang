package account

import(
	"math/rand"
	"time"
	"strconv"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto"
)
type LoginService struct{}
func (ac *LoginService) GetRandomSignMsg()string{
	rand.Seed(time.Now().UnixNano())
	//生成六位随机数
	nonce := int64(rand.Intn(1000000))
	msgString := "I confirm to sign with random nonce:"+strconv.FormatInt(nonce , 10)
	return msgString
}
/* func (ac *LoginService)Login(message string, signedMessage string) (string, error) {
    // Hash the unsigned message using EIP-191
    hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
    hash := crypto.Keccak256Hash(hashedMessage)

    // Get the bytes of the signed message
    decodedMessage := hexutil.MustDecode(signedMessage)

    // Handles cases where EIP-115 is not implemented (most wallets don't implement it)
    if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
        decodedMessage[64] -= 27
    }

    // Recover a public key from the signed message
    sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
    if sigPublicKeyECDSA == nil {
        err = errors.New("Could not get a public get from the message signature")
    }
    if err != nil {
        return "", err
    }

    return crypto.PubkeyToAddress(*sigPublicKeyECDSA).String(), nil
}
 */