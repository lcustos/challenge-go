package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	NameOfProgramm := os.Args
	for _, i := range NameOfProgramm[0] {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
