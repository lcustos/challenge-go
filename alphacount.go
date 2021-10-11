package tls_challenge_go_21_22

func AlphaCount(s string) int {
	count := 0
	for _, str := range s {
		for i := 'a'; i <= 'z'; i++ {
			if str == i {
				count++
			}
		}
		for j := 'A'; j <= 'Z'; j++ {
			if str == j {
				count++
			}
		}
	}
	return count
}
