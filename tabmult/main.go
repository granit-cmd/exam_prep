package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	num := args[0]

	n, _ := strconv.Atoi(num)

	ans1 := " x "
	ans2 := " = "

	for i := 1; i <= 9; i++ {
		tmp := i * n
		PrintRune(Itoa(i))
		PrintRune(ans1)
		PrintRune(num)
		PrintRune(ans2)
		PrintRune(Itoa(tmp))
		z01.PrintRune('\n')

	}
}

func PrintRune(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func Itoa(n int) string {
	ans := ""
	for n > 0 {
		ans = string(rune(n%10+'0')) + ans
		n /= 10
	}
	return ans
}

func Atoi(s string) (int, bool) {
	neg := 1
	if s[0] == '-' {
		neg = -1
		s = s[1:]
	}
	ans := 0

	for _, v := range s {
		if v >= '0' && v <= '9' {
			ans = ans*10 + int(v-'0')
		} else {
			return 0, false
		}
	}
	return ans * neg, true
}
