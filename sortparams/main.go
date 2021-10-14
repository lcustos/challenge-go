package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	for i := 1; i <= len(os.Args)-1; i++ {
		for j := i + 1; j <= len(os.Args)-1; j++ {
			if os.Args[i] > os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}

	for _, element := range os.Args[1:] {
		for _, word := range element {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
