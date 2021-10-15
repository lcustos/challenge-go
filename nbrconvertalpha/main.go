package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	flag := false
	for _, val := range arg {
		if val == "--upper" {
			flag = true
			arg = os.Args[2:]
		}
	}
	if flag {
		for i := 0; i <= len(arg)-1; i++ {
			z01.PrintRune(ToUpper(ConvertToLetter(arg[i])))
		}
		z01.PrintRune(10)
	} else if len(arg)-1 == 0 {
		z01.PrintRune(10)
	} else {
		for i := 0; i <= len(arg)-1; i++ {
			z01.PrintRune(ConvertToLetter(arg[i]))
		}
		z01.PrintRune(10)
	}
}

func ConvertToLetter(s string) rune {
	numb := 0
	for _, i := range s {
		if i == 0 {
			return 0
		}
		count := 0
		for k := '0'; k < i; k++ {
			count++
		}
		numb = numb*10 + count
	}
	numb = numb + 96
	if numb >= 123 {
		return ' '
	}
	return rune(numb)
}

func ToUpper(r rune) rune {
	if r-32 < 65 {
		return r
	}
	return r - 32
}
