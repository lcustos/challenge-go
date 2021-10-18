package tls_challenge_go_21_22

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	Range := make([]int, max-min)
	for i := 0; i < cap(Range); i++ {
		Range[i] = min + i
	}
	return Range
}
