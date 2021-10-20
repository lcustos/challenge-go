package tls_challenge_go_21_22

func SortWordArr(a []string) {
	i := 1
	for i < len(a) {
		if a[i-1] > a[i] {
			a[i-1], a[i] = a[i], a[i-1]
			i = 1
		} else {
			i++
		}
	}
}
