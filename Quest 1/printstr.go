package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, char := range s {
		_ = z01.PrintRune(char)
	}
}
