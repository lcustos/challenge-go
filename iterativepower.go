package tls_challenge_go_21_22

func IterativePower(nb int, power int) int {
	if nb > 25 {
		return 0
	}
	n := nb
	if power == 0 {
		return 0
	}
	for i := 1; i < power; i++ {
		nb = nb * n
	}
	return nb
}
