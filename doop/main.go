package main

import (
	"fmt"
	"os"
	"strconv"
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
		fmt.Print()
		return
	} else if validateOperator(args[1]) == false {
		fmt.Print()
		return
	}
	premier, _ := strconv.Atoi(args[0])
	second, _ := strconv.Atoi(args[2])
	if validNumber(premier, second) == true {
		if args[1] == "%" && second == 0 {
			fmt.Print("No Modulo by 0\n")
		} else if args[1] == "/" && second == 0 {
			fmt.Print("No division by 0\n")
		} else if args[1] == "+" {
			fmt.Println(premier + second)
		} else if args[1] == "-" {
			fmt.Println(premier - second)
		} else if args[1] == "*" {
			fmt.Println(premier * second)
		} else if args[1] == "/" {
			fmt.Println(premier / second)
		} else {
			fmt.Println(premier % second)
		}
	} else {
		fmt.Print()
		return
	}
}
