package tls_challenge_go_21_22

func AtoiBase(s string, base string) int {
	var res int
	if len(base) <= 1 {
		return 0
	}
	for _, word := range base {
		if string(word) == "-" {
			return 0
		}
	}
	for i := 0; i <= len(base)-1; i++ {
		for j := i + 1; j <= len(base)-1; j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	for k := range s {
		for l := range base {
			if s[k] == base[l] {
				for m := 0; m < (len(s)-1)-k; m++ {
					l = l * len(base)
				}
				res += l
			}
		}
	}
	return res
}
