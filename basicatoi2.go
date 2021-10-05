package tls_challenge_go_21_22

import "strconv"

func BasicAtoi2(s string) int {
	i, err := strconv.Atoi(s)

	if err == nil {
		return i
	} else {
		return 0
	}
}
