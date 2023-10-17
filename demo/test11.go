package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func main() {

	aa := 11525.29
	fmt.Println(float32(aa))
	s := strconv.FormatFloat(aa, 'f', -1, 64) //float64转string
	fmt.Println(s)

	// 对于保留小数的处理
	pi := decimal.NewFromFloat(3.1415926535897932384626)
	pi1 := pi.Round(3)    // 对pi值四舍五入保留3位小数
	fmt.Println(pi1)      // 3.142
	pi2 := pi.Truncate(3) // 对pi值保留3位小数之后直接舍弃
	fmt.Println(pi2)      // 3.141
}
