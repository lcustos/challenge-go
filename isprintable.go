package tls_challenge_go_21_22

func IsPrintable(s string) bool {
	IU := true
	for i := 0; i <= len(s)-1; i++ {
		if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' || rune(s[i]) >= 'a' && rune(s[i]) <= 'z' || rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			IU = true
		} else {
			return false
		}
	}
	return IU
}
