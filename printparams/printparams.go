package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, element := range os.Args[1:] {
		for _, word := range element {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
