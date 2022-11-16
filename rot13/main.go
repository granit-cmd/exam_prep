package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	for _, v := range os.Args[1] {
		if 'a' <= v && v <= 'n' || 'A' <= v && v <= 'N' {
			v += 13
		} else if 'm' <= v && v <= 'z' || 'M' <= v && v <= 'Z' {
			v -= 13
		}
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
