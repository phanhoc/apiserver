package api

import (
	"errors"
	"fmt"

	bip39 "github.com/tyler-smith/go-bip39"
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
