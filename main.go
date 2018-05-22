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
	seed, err := api.CreateSeed("globe chief envelope canyon first squeeze fold acid denial lawn beach moral slot extend endorse usage art mean faculty analyst pigeon nose split bullet")
	if err != nil {
		fmt.Println("CreateSeed error: ", err)
		return
	}
	fmt.Println("Seed:")
	s := string(seed)
	fmt.Println(s)
}
