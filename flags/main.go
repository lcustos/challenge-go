package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, word := range os.Args {
		if word == "--help" || word == "--h" || len(os.Args) == 0 {
		}
		if word == "--insert" || word == "--i" {
			for _, element := range os.Args[1:] {
				for _, word1 := range element {
					z01.PrintRune(word1)
				}
				z01.PrintRune('\n')
			}
		}
		if word == "--order" {
			for i := 1; i <= len(os.Args)-1; i++ {
				for j := i + 1; j <= len(os.Args)-1; j++ {
					if os.Args[i] > os.Args[j] {
						os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
					}
				}
			}
			for _, element := range os.Args[1:] {
				for _, word1 := range element {
					z01.PrintRune(word1)
				}
				z01.PrintRune('\n')
			}
		}
	}
}
