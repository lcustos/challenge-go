package tls_challenge_go_21_22

func RecursiveFactorial(nb int) int {
	if nb > 25 || nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	} else {
		return RecursiveFactorial(nb-1) * nb
	}
}
