package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	s := ""
	for i := 0; i < len(a); i++ {
		s = a[i]
		for _, word := range s {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
