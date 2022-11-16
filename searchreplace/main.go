package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	first := args[0]
	second := args[1]
	third := args[2]
	for i := 0; i <= len(first)-1; i++ {
		if string(first[i]) == second {
			first = first[:i] + third + string(first[i+1:])
		}
	}
	for _, v := range first {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
