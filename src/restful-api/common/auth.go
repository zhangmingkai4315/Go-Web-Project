package common

import (
	"io/ioutil"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatelf("[initKeys]:%s\n", err)
	}
	verifyKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatelf("[initKeys]:%s\n", err)
		panic(err)
	}

}
