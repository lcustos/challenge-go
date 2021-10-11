package tls_challenge_go_21_22

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if nb > 25 || power > 25 || power <= 0 {
		return 0
	}
	return nb * RecursivePower(nb, power-1)
}
