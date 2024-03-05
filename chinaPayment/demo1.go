package main

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func generateMerchantSignature(params map[string]interface{}, secret string) string {
	// 步骤1: 将参数按字段名排序
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 步骤2: 拼接非空参数
	var string1 string
	for _, key := range keys {
		if key != "sign" && params[key] != nil {
			if string1 != "" {
				string1 += "&"
			}
			string1 += fmt.Sprintf("%s=%v", key, params[key])
		}
	}

	// 步骤3: 在string1后追加&key={merchant secret}
	string2 := string1 + "&key=" + secret

	// 步骤4: 计算MD5哈希值并转换为大写
	hash := md5.New()
	hash.Write([]byte(string2))
	signature := fmt.Sprintf("%X", hash.Sum(nil))

	return signature
}

func main() {
	// 测试参数和密钥
	params := map[string]interface{}{}

	merchantSecret := "767D2FC0714C42289DB4"

	// 生成商家签名
	merchantSignature := generateMerchantSignature(params, merchantSecret)
	fmt.Println("Generated Merchant Signature:", merchantSignature)

}
