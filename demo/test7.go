package main

import (
	"log"
)

type demo struct {
	Num   int
	Value string
}

// 记录相邻的重复元素
func main() {
	input := []demo{{
		Num:   0,
		Value: "0",
	}, {
		Num:   1,
		Value: "1",
	}, {
		Num:   2,
		Value: "1",
	}, {
		Num:   3,
		Value: "1",
	}, {
		Num:   4,
		Value: "1",
	}, {
		Num:   5,
		Value: "5",
	}, {
		Num:   6,
		Value: "1",
	}, {
		Num:   7,
		Value: "1",
	}, {
		Num:   8,
		Value: "1",
	}}
	output := cleans(input)
	log.Println(output)
}
func cleans(input []demo) []demo {
	ress := make([][]demo, 0, len(input))
	res := make([]demo, 0, len(input))
	for i := 0; i < len(input); i++ {

		if i == len(input)-1 {
			if input[i].Value == input[i-1].Value {
				res = append(res, input[i])
			}
			if len(res) > 0 {
				ress = append(ress, duplicateRemoval(res))
				res = make([]demo, 0, len(input))
			}

			continue
		}
		if input[i].Value == input[i+1].Value {
			res = append(res, input[i])
			res = append(res, input[i+1])

		} else {
			if len(res) > 0 {
				ress = append(ress, duplicateRemoval(res))
				res = make([]demo, 0, len(input))
			}

		}
	}
	log.Println(ress)
	return res
}
func duplicateRemoval(input []demo) []demo {

	aa := make([]demo, 0, len(input))

	for i := range input {
		flag := true
		for j := range aa {
			if input[i].Num == aa[j].Num {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			aa = append(aa, input[i])
		}
	}
	return aa
}
