package main

import (
	"fmt"
	"github.com/agnivade/levenshtein"
)

// 字符串相似度
func main() {
	s1 := "sitt"
	s2 := "sitting"
	distance := levenshtein.ComputeDistance(s1, s2)
	fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
	// Output:
	// The distance between kitten and sitting is 3.
}
