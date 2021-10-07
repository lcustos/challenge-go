package tls_challenge_go_21_22

func Sqrt(nb int) int {
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
