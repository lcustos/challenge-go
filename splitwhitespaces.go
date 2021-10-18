package tls_challenge_go_21_22

func SplitWhiteSpaces(s string) []string {
	space := 0
	len2 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if s[i-1] == ' ' || s[i-1] == '\t' || s[i-1] == '\n' {
				space--
			}
			space++
		}
	}
	if len(s) == 0 {
		res := []string{""}
		return res
	}
	for i := range s {
		len2 = i
	}
	res := make([]string, space+1)
	check := true
	k := 0
	for i := 0; i <= len2; i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if !check {
				k++
			}
			check = true
			continue
		}
		res[k] += string(s[i])
		check = false
	}
	return res
}
