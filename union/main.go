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
	arg := os.Args[1] + os.Args[2]
	var res string
	res += string(arg[0])

	for i := 1; i < len(arg); i++ {
		k := 0
		for j := 0; j < len(res); j++ {

			if arg[i] == res[j] {
				k++
			}
			if j == len(res)-1 && k == 0 {
				res += string(arg[i])
			}
		}
	}
	Println(res)
}

func Println(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
