package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func calcComb(n int) []string {
	if n == 1 {
		return []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	}
	var precedent = calcComb(n - 1)
	var result []string
	for _, str := range precedent {
		for i := str[len(str)-1] - '0' + 1; i < 10; i++ {
			result = append(result, str+string(rune(i+'0')))
		}
	}
	return result
}

func printStr(str string) {
	for _, char := range str {
		_ = z01.PrintRune(char)
	}
}

func PrintCombN(n int) {
	result := calcComb(n)
	for i, str := range result {
		printStr(str)
		if i+1 != len(result) {
			_ = z01.PrintRune(',')
			_ = z01.PrintRune(' ')

		}
	}
	_ = z01.PrintRune('\n')
}
