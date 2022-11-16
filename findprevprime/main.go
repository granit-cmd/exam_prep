package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindPrevPrime(6))
	fmt.Println(FindPrevPrime(10))
}

func FindPrevPrime(nb int) int {
	if nb%2 != 0 {
		return nb
	} else {
		nb = nb - 1
		return nb
	}
}
