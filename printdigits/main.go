package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
	invalidRune := rune(-1)
	err := z01.PrintRune(invalidRune)
	if err == nil {
		panic("z01.PrintRune should fail with an invalid rune")
	}
}
