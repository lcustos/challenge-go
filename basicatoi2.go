package tls_challenge_go_21_22

func BasicAtoi2(s string) int {
	o_number := 0
	c := 0
	a_s := []rune(s)
	for _, word := range a_s {
		if word > 57 || word == 32 {
			return 0
		} else {
			for i := '0'; i < word; i++ {
				c++
			}
			o_number = o_number*10 + c
			c = 0
		}
	}
	return o_number
}
