package tls_challenge_go_21_22

func Any(f func(string) bool, a []string) bool {
	r := false
	for _, i := range a {
		if f(i) == true {
			return true
		} else {
			r = false
		}
	}
	return r
}
