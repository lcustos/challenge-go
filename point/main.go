package main

import (
	"github.com/01-edu/z01"
)

func setPoint(x *int, y *int) {
	*x = 42
	*y = 21
}

func main() {
	x := 0
	y := 0
	setPoint(&x, &y)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(y)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n/10 != 0 {
			PrintNbr(n / -10)
		}
		mod := '0'
		for i := 0; i < -(n % 10); i++ {
			mod++
		}
		z01.PrintRune(mod)
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		if n/10 != 0 {
			PrintNbr(n / 10)
		}
		mod := '0'
		for i := 0; i < n%10; i++ {
			mod++
		}
		z01.PrintRune(mod)
	}
}
