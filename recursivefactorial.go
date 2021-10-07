package tls_challenge_go_21_22

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
