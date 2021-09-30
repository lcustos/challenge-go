package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i-- {
		z01.PrintRune(i)
	}
	invalidRune := rune(-1)
	err := z01.PrintRune(invalidRune)
	if err == nil {
		panic("z01.PrintRune should fail with an invalid rune")
	}
}
