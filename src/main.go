package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	length := int32(len(s))
	countIfSumEquareTo_d := int32(0)

	for i := int32(0); i <= length-m; i++ {
		var sumOfEach int32
		for j := int32(0); j < m; j++ {
			sumOfEach += s[j+i]
		}
		if sumOfEach == d {
			countIfSumEquareTo_d++
		}
	}

	return countIfSumEquareTo_d
}

func main() {
	s := []int32{4}
	var d int32
	d = 4
	var m int32 = 1
	fmt.Println(birthday(s, d, m))
}
