package tls_challenge_go_21_22

import "strconv"

func BasicAtoi(s string) int {
	if s, err := strconv.Atoi(s); err == nil {
		return s
	}
	return 0
}
