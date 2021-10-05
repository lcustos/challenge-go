package main

import (
	piscine "challenge-go"
	"fmt"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
