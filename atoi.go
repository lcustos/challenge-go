package tls_challenge_go_21_22

func Atoi(s string) int {
	negate := false
	table := []rune(s)
	o_number := 0
	c := 0
	a_s := []rune(s)
	e := 0
	if len(table) != 0 {
		if table[0] == '-' {
			negate = true
		}
		for index, word := range a_s {
			if (word > '9' || word < '0') && index != 0 {
				return 0
			} else {
				for i := '0'; i < word; i++ {
					c++
				}
				o_number = o_number*10 + c
				c = 0
			}
		}
		if negate == false {
			e = o_number
			return e
		} else {
			e = o_number * -1
			return e
		}
	}
	return e
}
