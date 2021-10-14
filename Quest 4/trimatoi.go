package tls_challenge_go_21_22

func TrimAtoi(s string) int {
	negate := false
	table := []rune(s)
	result := 0
	c := 0
	if len(table) != 0 {
		for _, word := range table {
			if word < '0' || word > '9' {
				if word == '-' {
					if result == 0 {
						negate = true
					}
				}
			} else {
				for i := '0'; i < word; i++ {
					c++
				}
				result = result*10 + c
				c = 0
			}
		}
		if negate == false {
			return result
		} else {
			result *= -1
			return result
		}
	}
	return result
}
