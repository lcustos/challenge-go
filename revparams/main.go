package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	len := 0
	for i := range os.Args {
		len = i
	}
	for i := len; i >= 1; i-- {
		for _, k := range os.Args[i] {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
