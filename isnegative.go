package isnegative

import (
	"github.com/01-edu/z01"
	_ "github.com/01-edu/z01"
)

func isnegative(nb int) {
	if nb > 0 {
		z01.PrintRune('F')

	} else if nb < 0 {
		z01.PrintRune('T')

	}
	z01.PrintRune('\n')
}
