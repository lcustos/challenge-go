package tls_challenge_go_21_22

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return '\n'
}
