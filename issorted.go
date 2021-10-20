package tls_challenge_go_21_22

func IsSorted(f func(a, b int) int, a []int) bool {
	s := false
	if a[0] < a[1] {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) > 0 {
				s = true
			} else if f(a[i-1], a[i]) <= 0 {
			}
			return false
		}
	} else if a[0] > a[1] {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) < 0 {
				s = true
			} else if f(a[i-1], a[i]) >= 0 {
				return false
			}
		}
	} else {
		return false
	}
	return s
}
