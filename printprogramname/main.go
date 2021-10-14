package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var index int
	NameOfProgramm := os.Args[0]
	for i := 0; i < len(NameOfProgramm); i++ {
		if NameOfProgramm[i] == '\\' || NameOfProgramm[i] == '/' {
			index = i
		}
	}
	for _, element := range NameOfProgramm[index+1:] {
		if element == '.' {
			break
		}
		z01.PrintRune(element)
	}
	z01.PrintRune('\n')
}
