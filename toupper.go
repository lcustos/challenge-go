package tls_challenge_go_21_22

func ToUpper(s string) string {
	var t string
	for i := 0; i <= len(s)-1; i++ {
		if rune(s[i]) >= 'a' && rune(s[i]) <= 'z' {
			t = t + string(rune(s[i])-32)
		} else {
			t = t + string(s[i])
		}
	}
	return t
}
