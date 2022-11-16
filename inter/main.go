package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	var result string

	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(arg2); j++ {
			if arg1[i] == arg2[j] {

				if len(result) > 0 {
					result = repeatChecker(result, string(arg1[i]))
				} else {
					result += string(arg1[i])
				}

				break
			}
		}
	}

	Println(result)
	z01.PrintRune('\n')
}

func repeatChecker(result string, letter string) string {
	checker := 0
	for k := 0; k < len(result); k++ {
		if result[k] == letter[0] {
			checker++
		}
		if k == len(result)-1 && checker == 0 {
			result += letter
			break
		}
	}
	return result
}

func Println(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
