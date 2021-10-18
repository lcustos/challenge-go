package tls_challenge_go_21_22

func Split(s, sep string) []string {
	sepa := 0
	for j := 0; j <= len(s)-len(sep); j++ {
		if s[j:j+len(sep)] == sep {
			sepa++
		}
	}
	if len(s) == 0 {
		res := []string{""}
		return res
	}
	res := make([]string, sepa+1)
	check := true
	k := 0
	for j := 0; j <= len(s)-len(sep); j++ {
		if s[j:j+len(sep)] == sep {
			if !check {
				k++
			}
			check = true
			j++
			continue
		}
		res[k] += string(s[j])
		check = false
	}
	return res
}
