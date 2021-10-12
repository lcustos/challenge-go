package tls_challenge_go_21_22

func Join(strs []string, sep string) string {
	concat := ""
	for _, word := range strs {
		concat += word
		concat += sep
	}
	return concat
}
