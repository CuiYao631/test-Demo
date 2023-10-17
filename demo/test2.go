package main

import (
	"fmt"
	"time"
)

// 删除相邻的重复元素
func main() {
	//input := []string{"1", "2", "2", "3", "4", "5", "6", "5"}
	//output := clean(input)
	//log.Println(output)
	input02 := []string{"a", "b", "c", "b", "d", "a"}
	result02 := unUserIdRepeat02(input02)
	fmt.Println(result02) // [a b c d]
	input := []string{"a", "b", "c", "b", "d", "a"}
	result := unUserIdRepeat(input)
	fmt.Println(result) // [a b c d]

}

func clean(input []string) []string {

	res := make([]string, 0, len(input))
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 {
			res = append(res, input[i])
			continue
		}
		if input[i] != input[i+1] {
			res = append(res, input[i])
		}
	}
	return res
}
func unUserIdRepeat(input []string) []string {
	start := time.Now()
	encountered := map[string]bool{}
	result := make([]string, 0, len(input))

	for _, v := range input {
		if encountered[v] == true {
			// 如果当前元素已经出现过，直接跳过
			continue
		} else {
			// 如果当前元素没有出现过，加入结果数组并标记为已出现
			encountered[v] = true
			result = append(result, v)
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	return result
}
func unUserIdRepeat02(input []string) []string {
	start := time.Now()
	aa := make([]string, 0, len(input))
	for i := range input {
		flag := true
		for j := range aa {
			if input[i] == aa[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			aa = append(aa, input[i])
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	return aa
}
