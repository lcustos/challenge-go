package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args); i >= 0; i-- {
		for _, k := range os.Args[i] {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
