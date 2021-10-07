package tls_challenge_go_21_22

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		for j := 2; j < nb; j++ {
			if i*j == nb {
				return false
			}
		}
	}
	return true
}
