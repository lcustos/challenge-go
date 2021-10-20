package tls_challenge_go_21_22

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, i := range tab {
		if f(i) == true {
			count++
		}
	}
	return count
}
