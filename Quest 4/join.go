package tls_challenge_go_21_22

func Join(strs []string, sep string) string {
	concat := ""
	for i, word := range strs {
		concat += word
		i++
		if i < len(strs) {
			concat += sep
		}
	}
	return concat
}
