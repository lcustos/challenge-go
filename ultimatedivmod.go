package tls_challenge_go_21_22

func UltimateDivMod(a *int, b *int) {
	c := *a
	*a = *a / *b
	*b = c % *b

}
