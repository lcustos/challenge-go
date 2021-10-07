package tls_challenge_go_21_22

func StrRev(s string) string {
	lister := []rune(s)
	for i, j := 0, len(lister)-1; i < j; i, j = i+1, j-1 {
		lister[i], lister[j] = lister[j], lister[i]
	}
	return string(lister)
}
