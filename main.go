package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase0 := "password1"
	acc0, err := key.NewAccount(passphrase0)
	if err != nil {
		log.Error(err.Error())
	}

	fmt.Println(acc0.Address)
}
