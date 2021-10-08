package tls_challenge_go_21_22

func IsPrime(nb int) bool {
	if nb == 0 || nb == 1 {
		return false
	}
	if nb < 0 {
		nb = -nb
	}
	for i := 2; i <= nb*nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
