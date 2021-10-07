package tls_challenge_go_21_22

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		for i2 := 2; i2 < nb; i2++ {
			if i*i2 == nb {
				return false
			}
		}
	}
	return true
}
