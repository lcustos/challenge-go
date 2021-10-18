package tls_challenge_go_21_22

func SplitWhiteSpaces(s string) []string {
	len1 := 0
	len2 := 0
	for _, i := range s {
		if i == ' ' || i == '\t' || i == '\n' {
			len1++
		}
	}
	len1--
	if len1 < 0 {
		res := []string{""}
		return res
	}
	if len1 == 0 {
		len1++
	}
	for i := range s {
		len2 = i
	}
	res := make([]string, len1+1)
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
