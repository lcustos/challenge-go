package tls_challenge_go_21_22

func SplitWhiteSpaces(s string) []string {
	len := 0
	len2 := 0
	for _, i := range s {
		if i == ' ' || i == '\t' || i == '\n' {
			len++
		}
	}
	len++
	if s[0] == ' ' {
		len--
	}
	if s[len] == ' ' {
		len--
	}
	for i := range s {
		len2 = i
	}

	res := make([]string, len)
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
