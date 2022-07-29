package main

import (
	"fmt"
	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/sr25519"
	"log"
)

func main() {

}

// seed 解出所有
func seedToAll(seed string) {
	scheme := sr25519.Scheme{}
	keyPair, err := subkey.DeriveKeyPair(scheme, seed)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("publicKey: %x\n", keyPair.Public())
	fmt.Printf("accountid: %x\n", keyPair.AccountID()) // public key 和 account id 是一样的
	fmt.Printf("address: %s\n", keyPair.SS58Address(42))
}

// 从 ss58 地址可以解出公钥
func ss58ToPublicKey(ss58Address string) []byte {
	format, publicKey, err := subkey.SS58Decode(ss58Address)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("network: %v\npublic key hex: 0x%x\n", format, publicKey)
	return publicKey
}

// 使用 seed 签名
func sign(seed string, data []byte) []byte {
	scheme := sr25519.Scheme{}
	keyPair, err := subkey.DeriveKeyPair(scheme, seed)
	if err != nil {
		log.Fatalln(err)
	}
	signData, err := keyPair.Sign(data)
	if err != nil {
		log.Fatalln(err)
	}
	return signData
}

// 只要知道 ss58 地址，就可以验证数据
func verify(ss58 string, data []byte, sign []byte) bool {
	publicKey := ss58ToPublicKey(ss58)
	scheme := sr25519.Scheme{}
	key, err := scheme.FromPublicKey(publicKey)
	if err != nil {
		log.Fatalln(err)
	}
	return key.Verify(data, sign)
}
