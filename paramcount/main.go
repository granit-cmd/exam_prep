package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	str := "0123456789"
	strlen := len(arg)
	z01.PrintRune(rune(str[strlen]))
	z01.PrintRune('\n')
}
