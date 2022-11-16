package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	runes := []rune(args[0])
	for _, v := range runes {
		if v >= 'a' && v <= 'z' {
			v = 'z' - v + 'a'
		}
		z01.PrintRune(v)

	}
	z01.PrintRune('\n')
}
