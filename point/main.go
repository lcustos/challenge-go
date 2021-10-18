package main

import "fmt"

func setPoint(x *int, y *int) {
	*x = 42
	*y = 21
}

func main() {
	x := 0
	y := 0
	setPoint(&x, &y)
	fmt.Printf("x = %d, y = %d", x, y)
}
