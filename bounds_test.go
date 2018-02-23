package boundsChecking

import "testing"

const v = 0x0807060504030201

var resultBoundsCheck []byte
var resultNoBoundsCheck []byte

func BenchmarkBoundsCheck(b *testing.B) {
	var r []byte
	for n := 0; n < b.N; n++ {
		r = BoundsCheck(v)
	}
	resultBoundsCheck = r
}

func BenchmarkNoBoundsCheck(b *testing.B) {
	var r []byte
	for n := 0; n < b.N; n++ {
		r = NoBoundsCheck(v)
	}
	resultNoBoundsCheck = r
}
