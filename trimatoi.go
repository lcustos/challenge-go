package tls_challenge_go_21_22

func TrimAtoi(s string) int {
	negate := false
	table := []rune(s)
	oNumber := 0
	c := 0
	e := 0
	if len(table) != 0 {
		for _, word := range table {
			if word > '9' || word < '0' {
				if word == '-' {
					if oNumber == 0 {
						negate = true
					}
				}
			} else {
				for i := '0'; i < word; i++ {
					c++
				}
				oNumber = oNumber*10 + c
				c = 0
			}
		}
		if negate == false {
			e = oNumber
			return e
		} else {
			e = oNumber * -1
			return e
		}
	}
	return e
}
