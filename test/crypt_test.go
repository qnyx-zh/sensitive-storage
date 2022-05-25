package test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"sensitive-storage/util/crypt"
	"testing"
)

func Test_crypt(t *testing.T) {
	d := []byte("asdfghjkloiuytre")
	key := []byte("hgfedcba87654321")
	fmt.Println("加密前:", string(d))
	x1, err := encryptAES(d, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("加密后:", base64.StdEncoding.EncodeToString(x1))
	x2, err := decryptAES(x1, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解密后:", string(x2))
	fmt.Printf("crypt.Md5crypt(\"zzz\"): %v\n", crypt.Md5crypt("12345"))
	fmt.Printf("crypt.Md5crypt(\"zzz\"): %v\n", crypt.Md5crypt("zzz"))
	s, _ := crypt.AesEncrypt("adasfas")
	fmt.Printf("s: %v\n", s)
	s2, _ := crypt.AesDeCrypt(s)
	fmt.Printf("s2: %v\n", s2)
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
