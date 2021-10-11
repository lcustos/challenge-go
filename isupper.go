package tls_challenge_go_21_22

func IsUpper(s string) bool {
	IU := true
	for i := 0; i <= len(s)-1; i++ {
		if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
			IU = true
		} else {
			IU = false
		}
	}
	return IU
}
