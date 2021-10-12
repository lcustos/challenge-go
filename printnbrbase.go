package tls_challenge_go_21_22

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(s int, t string) {
	var a int
	tab := []int{}
	if len(t) <= 1 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for _, word := range t {
		if string(word) == "-" {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	for i := 0; i <= len(t)-1; i++ {
		for j := i + 1; j <= len(t)-1; j++ {
			if t[i] == t[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if s < 0 {
		z01.PrintRune('-')
		s = -s
	}
	for s > 0 {
		a = s % len(t)
		tab = append(tab, a)
		s /= len(t)
	}
	mod := 0
	for m := len(tab) - 1; m >= 0; m-- {
		for k := 0; k < int(t[tab[m]]); k++ {
			mod++
		}
		z01.PrintRune(rune(mod))
		mod = 0
	}
}
