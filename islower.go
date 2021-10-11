package tls_challenge_go_21_22

func IsLower(s string) bool {
	IU := true
	for i := 0; i <= len(s)-1; i++ {
		if rune(s[i]) >= 'a' && rune(s[i]) <= 'z' {
			IU = true
		} else {
			return false
		}
	}
	return IU
}
