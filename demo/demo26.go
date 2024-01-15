package main

import (
	"fmt"
	"strings"
)

func replaceCharacter(input, target, replacement string) string {
	// 使用 strings.Replace 函数进行替换
	result := strings.Replace(input, target, replacement, -1)
	return result
}

func main() {
	inputString := "2022第四季度员工工资发放记录"
	targetCharacter := "202qq2"
	replacementCharacter := "2023"

	// 调用替换函数
	result := replaceCharacter(inputString, targetCharacter, replacementCharacter)

	// 输出结果
	fmt.Println("Original String:", inputString)
	fmt.Println("Result after replacement:", result)
}
