package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	arg := []rune(args[0])
	// arg := "lorem,ipsum"
	str := ""
	j := 0
	flag := false

	for _, v := range arg {
		if v == ' ' {
			flag = true
			break
		}
	}
	if !flag {
		for _, v := range arg {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return
	}

loop:
	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] != ' ' {
			for j = i; j >= 0; j-- {
				if arg[j] != ' ' {
					str = string(arg[j]) + str
				} else {
					break loop
				}
			}
		} else if arg[i] == ' ' {
			continue
		}
	}

	for _, v := range str {
		z01.PrintRune(v)
	}

	if str == "" {
		return
	}
	z01.PrintRune('\n')
}
