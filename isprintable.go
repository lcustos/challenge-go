package tls_challenge_go_21_22

func IsPrintable(s string) bool {
	IU := true
	for i := 0; i <= len(s)-1; i++ {
		if rune(s[i]) >= 32 && rune(s[i]) <= 126 {
			IU = true
		} else {
			return false
		}
	}
	return IU
}
