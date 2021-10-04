package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _, char := range str {
		_ = z01.PrintRune(char)
	}
}
