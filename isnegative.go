package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune('F')

	} else if nb < 0 {
		z01.PrintRune('T')

	}
	z01.PrintRune('\n')
}
