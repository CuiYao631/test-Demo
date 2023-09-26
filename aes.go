package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
//key不能泄露
type Algorithm struct {
	key []byte
}

func NewAlgorithm() *Algorithm {
	return &Algorithm{key: []byte("1234asdf1234asdf")}
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// EncryptByAes Aes加密 后 base64
func (a *Algorithm) EncryptByAes(data string) (string, error) {

	res, err := AesEncrypt([]byte(data), a.key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

// DecryptByAes base64解码后 Aes 解密
func (a *Algorithm) DecryptByAes(data string) (string, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	res, err := AesDecrypt(dataByte, a.key)
	return string(res), nil
}

//加密手机号
func (a *Algorithm) encryptPhone(input string) string {
	if len([]rune(input)) != 11 {
		log.Println("电话长度不对")
	}
	s1, _ := a.EncryptByAes(input[0:6])
	s2, _ := a.EncryptByAes(input[6:13])
	s3, _ := a.EncryptByAes(input[13:18])
	out := s1 + "@" + s2 + "@" + s3

	return out
}

//解密手机号
func (a *Algorithm) decryptPhone(input string) string {
	res := strings.Split(input, "@")
	e := ""
	for _, v := range res {
		decrypt, _ := a.DecryptByAes(v)
		e += decrypt
	}

	return e
}

//0x20~0x7E之间的可见字符
func main() {
	a := NewAlgorithm()
	//加密
	//encrypt, _ := a.EncryptByAes("123456")
	//解密
	//decrypt, _ := a.DecryptByAes(encrypt)

	//打印
	//fmt.Printf("加密前: %s\n", "123456")
	//fmt.Printf("加密后: %s\n", encrypt)
	//fmt.Printf("解密后：%s\n", decrypt)
	aes := a.encryptPhone("610125199608210012")
	res := a.decryptPhone(aes)
	fmt.Printf("加密前: %s\n", "610125199608210012")
	fmt.Printf("加密后: %s\n", aes)
	fmt.Printf("解密后：%s\n", res)

}
