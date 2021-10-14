package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	NameOfProgramm := os.Args
	for _, i := range NameOfProgramm[0] {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
