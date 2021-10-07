package tls_challenge_go_21_22

func FindNextPrime(nb int) int {
	for i := nb; i <= nb*nb; i++ {
		if IsPrime(i) == true {
			return i
		}
	}
	return nb
}
