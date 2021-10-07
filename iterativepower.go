package tls_challenge_go_21_22

func IterativePower(nb int, power int) int {
	n := nb
	if power == 0 {
		return 1
	} else {
		for i := 1; i < power; i++ {
			n = nb * n
		}
		return n
	}
}
