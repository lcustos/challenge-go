package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "quest8.txt"
	if len(os.Args) < 2 {
		fmt.Print("File name missing")
		z01.PrintRune('\n')
		return
	}
	if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
		z01.PrintRune('\n')
		return
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(string(data))
	z01.PrintRune('\n')
}
