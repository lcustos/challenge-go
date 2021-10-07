package tls_challenge_go_21_22

func IsPrime(nb int) bool {
	if nb > 25 || nb == 0 {
		return false
	}
	if nb == 1 {
		return false
	}
	if nb < 0 {
		nb = nb * -1
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
