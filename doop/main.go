package main

import (
	"os"
)

func validateOperator(test string) bool {
	op := []string{"+", "-", "*", "/", "%"}
	for _, res := range op {
		if res == test {
			return true
		}
	}
	return false
}

func validNumber(a, b int) bool {
	if a <= -9223372036854775808 || a >= 9223372036854775807 || b <= -9223372036854775808 || b >= 9223372036854775807 {
		return false
	} else {
		return true
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 3 || len(args) < 3 {
		os.Stderr.WriteString(" ")
		return
	} else if validateOperator(args[1]) == false {
		os.Stderr.WriteString("")
		return
	}
	premier := Atoi(args[0])
	second := Atoi(args[2])
	if validNumber(premier, second) == true {
		if args[1] == "%" && second == 0 {
			os.Stderr.WriteString("No Modulo by 0\n")
		} else if args[1] == "/" && second == 0 {
			os.Stderr.WriteString("No division by 0\n")
		} else if args[1] == "+" {
			os.Stderr.WriteString(string(premier + second))
		} else if args[1] == "-" {
			os.Stderr.WriteString(string(premier - second))
		} else if args[1] == "*" {
			os.Stderr.WriteString(string(premier * second))
		} else if args[1] == "/" {
			os.Stderr.WriteString(string(premier / second))
		} else {
			os.Stderr.WriteString(string(premier % second))
		}
	} else {
		os.Stderr.WriteString("")
		return
	}
}

func Atoi(s string) int {
	negate := false
	table := []rune(s)
	o_number := 0
	c := 0
	a_s := []rune(s)
	if len(table) != 0 {
		if table[0] == '-' {
			negate = true
		}
		for index, word := range a_s {
			if (word < '0' || word > '9') && index != 0 {
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
			return o_number
		} else {
			o_number *= -1
			return o_number
		}
	}
	return o_number
}
