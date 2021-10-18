package tls_challenge_go_21_22

func AppendRange(min, max int) []int {
	Range := []int(nil)
	if min >= max {
		return Range
	} else {
		for i := min; i < max; i++ {
			Range = append(Range, i)
		}
	}
	return Range
}
