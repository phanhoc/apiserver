package api

import (
	"errors"
	"fmt"
	"crypto/md5"
    "encoding/hex"

	bip39 "github.com/tyler-smith/go-bip39"
	aes "apiserver/utilities"
)

func CreateNewMnemonic() (string, error) {
	// Initial entropy with 256 bits
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		fmt.Println("Initial Entropy error")
		return "", err
	}
	// Generate new Mnemonic
	mnemonic, _ := bip39.NewMnemonic(entropy)
	if err != nil {
		fmt.Println("Generate Mnemonic error")
		return "", err
	}
	return mnemonic, err
}

func CreateSeed(mnemonic string) ([]byte, error) {
	// validate input
	if len(mnemonic) <= 0 {
		return nil, errors.New("Mnemonic unable to null or empty")
	}
	// generate seed from mnemonic
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	return seed, err
}

func CreateSeedHexString(mnemonic string) (string, error) {
	// validate input
	if len(mnemonic) <= 0 {
		return "", errors.New("Mnemonic unable to null or empty")
	}
	// generate seed from mnemonic
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	strSeed := bytesToHexString(seed)
	return strSeed, err
}

func EncodeSeed(key string, seed string)(string, error){
	// validate input
	if len(key) <=0 {
		return "", errors.New("Input invalidate")
	}
	if len(seed) <=0 {
		return "", errors.New("Input invalidate")
	}
	md5Key := getMD5Hash(key)
	byteKey := []byte(md5Key)
	return aes.Encrypt(byteKey, seed)
}

func DecodeSeed(key string, seed string)(string, error){
	// validate input
	if len(key) <=0 {
		return "", errors.New("Input invalidate")
	}
	if len(seed) <=0 {
		return "", errors.New("Input invalidate")
	}
	md5Key := getMD5Hash(key)
	byteKey := []byte(md5Key)
	return aes.Decrypt(byteKey, seed)
}

// BytesToHexString converts a byte slice into a hex string
func bytesToHexString(buf []byte) string {
	hlist := make([]byte, 2*len(buf))

	for i := range buf {
		hex := fmt.Sprintf("%02x", buf[i])
		idx := i * 2
		copy(hlist[idx:], hex)
	}
	return string(hlist)
}

func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}