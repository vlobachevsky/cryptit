package main

import (
	"fmt"
	"github.com/vlobachevsky/cryptit/encrypt"
	"github.com/vlobachevsky/cryptit/decrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Vitaly") 
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}