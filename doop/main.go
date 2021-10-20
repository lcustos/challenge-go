package main

import (
	tls_challenge_go_21_22 "challenge-go"
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
	premier := tls_challenge_go_21_22.BasicAtoi(args[0])
	second := tls_challenge_go_21_22.BasicAtoi(args[2])
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
