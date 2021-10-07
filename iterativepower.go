package tls_challenge_go_21_22

func IterativePower(nb int, power int) int {
	if nb > 25 || power > 25 || power <= 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	n := nb
	for i := 1; i < power; i++ {
		nb = nb * n
	}
	return nb
}