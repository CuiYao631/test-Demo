package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 生成json文件
func main() {

	mapInstances := make(map[string]string)
	mapInstances["a"] = "a"
	mapInstances["b"] = "b"

	// 创建文件
	filePtr, err := os.Create("DataForEdit.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(mapInstances)
	if err != nil {
		fmt.Println("生成文件错误", err.Error())
	} else {
		fmt.Println("生成文件成功")
	}
}
