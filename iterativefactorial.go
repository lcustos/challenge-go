package tls_challenge_go_21_22

func IterativeFactorial(nb int) int {
	f := 1
	if nb < 0 || nb >= 9223372036854775807 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			f = f * i
		}
		return f
	}
}