package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "quest8.txt"
	if len(os.Args) < 2 {
		fmt.Print("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
		return
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(string(data))
}
