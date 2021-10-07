package tls_challenge_go_21_22

import "github.com/01-edu/z01"

const N = '8'

var position = [N]rune{}

func isSafe(queen_number, row_position rune) bool {
	for i := '0'; i < queen_number; i++ {
		other_row_pos := position[i]

		if other_row_pos == row_position || other_row_pos == row_position-(queen_number-i) || other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

func solve(k rune) {
	if k == N {
		for i := '0'; i < N; i++ {
			z01.PrintRune(position[i] + 1)
		}
		z01.PrintRune('\n')
	} else {
		for i := '0'; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				solve(k + 1)
			}
		}
	}
}

func EightQueens() {
	solve('0')
}
