package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	p := "Bob hit a ball, the hit BALL flew far after it was hit."
	b := []string{"hit"}

	word := mostCommonWord(p, b)

	assert.Equal(t, "ball", word)
}

func Test2(t *testing.T) {
	p := "a."
	var b []string

	word := mostCommonWord(p, b)

	assert.Equal(t, "a", word)
}

func Test3(t *testing.T) {
	p := "Bob. hIt, baLl"
	b := []string{"bob", "hit"}

	word := mostCommonWord(p, b)

	assert.Equal(t, "ball", word)
}

func Test4(t *testing.T) {
	p := "a, a, a, a, b,b,b,c, c"
	b := []string{"a"}

	word := mostCommonWord(p, b)

	assert.Equal(t, "b", word)
}

func Test5(t *testing.T) {
	p := "abc abc? abcd the jeff!"
	b := []string{"abc", "abcd", "jeff"}

	word := mostCommonWord(p, b)

	assert.Equal(t, "the", word)
}

func Test6(t *testing.T) {
	p := "j. t? T. z! R, v, F' x! L; l! W. M; S. y? r! n; O. q; I? h; w. t; y; X? y, p. k! k, h, J, r? w! U! V; j' u; R! z. s. T' k. P? M' I' j! y. P, T! e; X. w? M! Y, X; G; d, X? S' F, K? V, r' v, v, D, w, K! S? Q! N. n. V. v. t? t' x! u. j; m; n! F, V' Y! h; c! V, v, X' X' t? n; N' r; x. W' P? W; p' q, S' X, J; R. x; z; z! G, U; m. P; o. P! Y; I, I' l' J? h; Q; s? U, q, x. J, T! o. z, N, L; u, w! u, S. Y! V; S? y' E! O; p' X, w. p' M, h! R; t? K? Y' z? T? w; u. q' R, q, T. R? I. R! t, X, s? u; z. u, Y, n' U; m; p? g' P? y' v, o? K? R. Q? I! c, X, x. r' u! m' y. t. W; x! K? B. v; m, k; k' x; Z! U! p. U? Q, t, u' E' n? S' w. y; W, x? r. p! Y? q, Y. t, Z' V, S. q; W. Z, z? x! k, I. n; x? z; V? s! g, U; E' m! Z? y' x? V! t, F. Z? Y' S! z, Y' T? x? v? o! l; d; G' L. L, Z? q. w' r? U! E, H. C, Q! O? w! s? w' D. R, Y? u. w, N. Z? h. M? o, B, g, Z! t! l, W? z, o? z, q! O? u, N; o' o? V; S! z; q! q. o, t! q! w! Z? Z? w, F? O' N' U' p? r' J' L; S. M; g' V. i, P, v, v, f; W? L, y! i' z; L? w. v, s! P?"
	b := []string{"m", "q", "e", "l", "c", "i", "z", "j", "g", "t", "w", "v", "h", "p", "d", "b", "a", "r", "x", "n"}

	word := mostCommonWord(p, b)

	assert.Equal(t, "y", word)
}
