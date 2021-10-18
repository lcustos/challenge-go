package tls_challenge_go_21_22

func ConcatParams(args []string) string {
	ans := ""
	count := 0
	for _, word := range args {
		if count < len(args)-1 {
			ans += word + "\n"
			count++
		} else {
			ans += word
		}
	}
	return ans
}
