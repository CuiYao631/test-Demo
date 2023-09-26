package main

import (
	"log"
	"regexp"
	"strings"
)

func main() {
	log.Println(KeyMatch3("/proxy/myid/res/res2", "/proxy/{id}/*"))
	log.Println(KeyMatch3("/proxy/myid", "/proxy/{id}/*"))
	log.Println(KeyMatch3("/proxy/myid/res", "/proxy/{id}/*"))
	log.Println(KeyMatch3("/proxy/rew/res", "/proxy/{id}/*"))
}

func KeyMatch3(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`\{[^/]+\}`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}
func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}

//func validateVariadicArgs(expectedLen int, args ...interface{}) error {
//	if len(args) != expectedLen {
//		return fmt.Errorf("Expected %d arguments, but got %d", expectedLen, len(args))
//	}
//
//	for _, p := range args {
//		_, ok := p.(string)
//		if !ok {
//			log.Println("Argument must be a string")
//		}
//	}
//
//	return nil
//}
//func KeyGet3Func(args ...interface{}) (interface{}, error) {
//	if err := validateVariadicArgs(3, args...); err != nil {
//		return false, fmt.Errorf("%s: %s", "keyGet3", err)
//	}
//
//	name1 := args[0].(string)
//	name2 := args[1].(string)
//	key := args[2].(string)
//
//	return KeyGet3(name1, name2, key), nil
//}
//func KeyGet3(key1, key2 string, pathVar string) string {
//	key2 = strings.Replace(key2, "/*", "/.*", -1)
//
//	re := regexp.MustCompile(`\{[^/]+?\}`) // non-greedy match of `{...}` to support multiple {} in `/.../`
//	keys := re.FindAllString(key2, -1)
//	key2 = re.ReplaceAllString(key2, "$1([^/]+?)$2")
//	key2 = "^" + key2 + "$"
//	re2 := regexp.MustCompile(key2)
//	values := re2.FindAllStringSubmatch(key1, -1)
//	if len(values) == 0 {
//		return ""
//	}
//	for i, key := range keys {
//		if pathVar == key[1:len(key)-1] {
//			return values[0][i+1]
//		}
//	}
//	return ""
//}
