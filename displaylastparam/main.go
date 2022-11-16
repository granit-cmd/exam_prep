package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		return
	}
	for _, v := range arg[len(arg)-1] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
