package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		a := string(bytes)
		io.WriteString(os.Stdout, a)
	} else {
		for i := 1; i < len(args); i++ {
			content, err := ioutil.ReadFile(args[i])
			strContent := string(content)
			if err != nil {
				errFirstPart := "ERROR: open "
				errScnPart := ": no such file or directory\n"
				finalErr := errFirstPart + args[i] + errScnPart
				io.WriteString(os.Stdout, finalErr)
				os.Exit(1)
			}
			for _, j := range strContent {
				z01.PrintRune(j)
			}
		}
	}
}
