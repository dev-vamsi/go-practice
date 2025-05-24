package main

import "math/rand"

var Letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890_/-")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(b)
}
