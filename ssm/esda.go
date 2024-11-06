package ssm

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func CreateESDAKeyPair() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("generate key fail", err)
		return "0x00", "0x00", "0x00", err
	}
	priKeyStr := hex.EncodeToString(crypto.FromECDSA(privateKey))
	pubKeyStr := hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey))
	decPubKey := hex.EncodeToString(crypto.CompressPubkey(&privateKey.PublicKey))

	return priKeyStr, pubKeyStr, decPubKey, nil
}

func SignMessage(privateKey string, txMsg string) (string, error) {
	hash := common.HexToHash(txMsg)
	privateByte, err := hex.DecodeString(privateKey)
	if err != nil {
		return "0x00", err
	}
	privateKeyEcdsa, err := crypto.ToECDSA(privateByte)
	if err != nil {
		return "0x00", err
	}
	signature, err := crypto.Sign(hash[:], privateKeyEcdsa)
	if err != nil {
		log.Error("sign transaction fail", "err", err)
		return "0x00", err
	}

	return hex.EncodeToString(signature), nil
}

func VerifySign(pubkey, msghash, signature string) bool {
	// pubkey []byte, digestHash []byte, signature []byte

	publicKey, err := hex.DecodeString(pubkey)
	if err != nil {
		return false
	}
	digestHash, err := hex.DecodeString(msghash)
	if err != nil {
		return false
	}
	signatureStr, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	return crypto.VerifySignature(publicKey, digestHash, signatureStr)
}
