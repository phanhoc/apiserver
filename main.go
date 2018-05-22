package main

import (
	api "apiserver/api"
	"fmt"
)

func main() {
	strMnemonic, err := api.CreateNewMnemonic()
	if err != nil {
		fmt.Println("CreateNewMnemonic error!")
		return
	}
	fmt.Println("strMnemonic:")
	fmt.Println(strMnemonic)
	seed, err := api.CreateSeedHexString("either hill rotate buzz echo pluck misery arm gain injury breeze symbol nation material pottery cradle hospital border creek harsh stool garage anxiety uniform")
	if err != nil {
		fmt.Println("CreateSeed error: ", err)
		return
	}
	fmt.Println("Seed:")
	fmt.Println(seed)
	key := "phanhoc@123"
	encodeSeed, err := api.EncodeSeed(key, seed)
	fmt.Println("Seed Encrytion:")
	fmt.Println(encodeSeed)
	decodeSeed, err := api.DecodeSeed(key, encodeSeed)
	fmt.Println("Seed Decrytion:")
	fmt.Println(decodeSeed)
}
