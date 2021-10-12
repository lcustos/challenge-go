package tls_challenge_go_21_22

import "fmt"

func BasicJoin(elems []string) string {
	concat := ""
	for _, word := range elems {
		fmt.Println(word)
		concat += word
	}
	return concat
}
