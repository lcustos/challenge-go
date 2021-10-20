package tls_challenge_go_21_22

func IsSorted(f func(a, b int) int, a []int) bool {
	verif := []bool{}
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			verif = append(verif, true)
		} else if f(a[i-1], a[i]) > 0 {
			verif = append(verif, false)
		} else {
			return false
		}
	}
	s := false
	for j := 1; j < len(verif); j++ {
		if verif[j-1] != verif[j] {
			return false
		} else {
			s = true
		}
	}
	return s
}
