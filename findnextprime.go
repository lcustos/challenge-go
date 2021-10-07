package tls_challenge_go_21_22

func FindNextPrime(nb int) int {
	for i := nb; i <= nb*nb; i++ {
		if i < 2 {
			i = 2
		}
		if IsPrime(i) == true {
			return i
		}
	}
	return nb
}
