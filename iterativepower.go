package tls_challenge_go_21_22

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if nb > 25 || power > 25 || power <= 0 {
		return 0
	}
	n := nb
	for i := 1; i < power; i++ {
		nb = nb * n
	}
	return nb
}
