package tls_challenge_go_21_22

func Capitalize(s string) string {
	var t string
	t = t + ToUpper(string(s[0]))
	for i := 1; i <= len(s)-1; i++ {
		if rune(s[i]) >= 'a' && rune(s[i]) <= 'z' || rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' || rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			if rune(s[i-1]) < '0' || rune(s[i-1]) > '9' && rune(s[i-1]) < 'A' || rune(s[i-1]) > 'Z' && rune(s[i-1]) < 'a' || rune(s[i-1]) > 'z' {
				t = t + ToUpper(string(s[i]))
			} else {
				t = t + ToLower(string(s[i]))
			}
		} else {
			t = t + string(s[i])
		}
	}
	return t
}
