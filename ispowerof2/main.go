package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	if len(arr) != 1 {
		return
	}

	n, _ := strconv.Atoi(arr[0])

	for n > 2 {
		if n%2 == 1 {
			Println("false")
			return
		}
		n /= 2
	}
	Println("true")
}

func Println(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
