package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {
	for _, element := range os.Args[1:] {
		for _, word := range element {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
