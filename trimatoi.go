package tls_challenge_go_21_22

func TrimAtoi(s string) int {
	negate := false
	table := []rune(s)
	o_number := 0
	c := 0
	a_s := []rune(s)
	e := 0
	if len(table) != 0 {
		for _, word := range a_s {
			if word > '9' || word < '0' {
				if word == '-' {
					if o_number == 0 {
						negate = true
					}
				}
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
