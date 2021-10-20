package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		mod := '0'
		for _, i := range SortIntegerTable2(ToMassive(n)) {
			for k := 0; k < i; k++ {
				mod++
			}
			z01.PrintRune(mod)
			mod = '0'
		}
	}
}

func ToMassive(n int) []int {
	var res []int
	for n >= 0 {
		if n == 0 {
			res = append(res, n)
		} else {
			a := n % 10
			res = append(res, a)
		}
		n = n / 10
	}
	return res
}

func SortIntegerTable2(table []int) []int {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			table[i], table[i-1] = table[i-1], table[i]
			i = 1
		} else {
			i++
		}
	}
	return table
}
