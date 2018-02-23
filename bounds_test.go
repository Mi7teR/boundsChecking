package boundsChecking

import "testing"

const v = 0x0807060504030201
func BenchmarkBoundsCheck(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BoundsCheck(v)
	}
}

func BenchmarkNoBoundsCheck(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NoBoundsCheck(v)
	}
}
