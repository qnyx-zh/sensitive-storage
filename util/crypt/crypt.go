package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
)

const (
	KEY = "asdfghjklqwertyu"
)

func AesEncrypt(src string) (string, error) {
	srcNew := []byte(src)
	b, err := encryptAES(srcNew, []byte(KEY))
	if err != nil {
		log.Printf("加密错误,原因=%v", err)
	}
	return base64.StdEncoding.EncodeToString(b), err
}
func AesDeCrypt(src string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(src)
	b, err := decryptAES(decode, []byte(KEY))
	return string(b), err
}

func Md5crypt(src string) string {
	b := md5.Sum([]byte(src))
	md5str := fmt.Sprintf("%x", b)
	return md5str
}

// 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 解密
func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}
