package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	seed          = "0x17403b2287de48c43934533f457f17f7cec505d9a54045567a9d121c3feb7b2e"
	publicKeyHex  = "0xfe664ccc5cd5fb02b96db97858642d4522b72b35876ca07d68d33af115a8175e"
	publicKeySS58 = "5HpGQhD72vZGgAFMMiCDY61mHYtANs6B4kZXrpptGm276KnT"
)

func TestVerify(t *testing.T) {

	// 要验证的数据
	data := []byte("hello!")

	// 前端client对数据进行签名
	signData := sign(seed, data)

	// 前端把签名后的数据和ss58地址传给后端，后端就可以验证签名
	verifyResult := verify(publicKeySS58, data, signData)

	assert.Equal(t, true, verifyResult)
}
