package tls_challenge_go_21_22

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}
