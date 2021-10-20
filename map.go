package tls_challenge_go_21_22

func Map(f func(int) bool, a []int) []bool {
	tab := []bool{}
	for _, i := range a {
		tab = append(tab, f(i))
	}
	return tab
}
