package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	res := a[0]
	for i := 0; i < len(a); i++ {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
