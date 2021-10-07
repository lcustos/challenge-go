package tls_challenge_go_21_22

func IterativeFactorial(nb int) int {
	f := 1
	if nb <= -4095264551424239652 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		f = f * i
	}
	return f
}
