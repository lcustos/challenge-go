package tls_challenge_go_21_22

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	} else {
		frune := []rune(s)
		return frune[n-1]
	}
}
