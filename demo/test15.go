package main

import (
	"fmt"
	"regexp"
	"strings"
)

var matchNonAlphaNumeric = regexp.MustCompile(`[^a-zA-Z0-9]+`)
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	str = matchNonAlphaNumeric.ReplaceAllString(str, "_")     //非常规字符转化为 _
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}") //拆分出连续大写
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")  //拆分单词
	return strings.ToLower(snake)                             //全部转小写
}

func main() {
	fmt.Println(ToSnakeCase("IamXXX"))
}
