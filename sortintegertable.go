package tls_challenge_go_21_22


func SortIntegerTable(table []int) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			table[i], table[i-1] = table[i-1], table[i]
			i = 1
		} else {
			i++
		}
	}
}
