package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := '0'; j <= '8'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i < j {
					if j < k {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}

				}

			}

		}

	}
}
