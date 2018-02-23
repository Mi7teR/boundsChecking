package boundsChecking

func BoundsCheck(v uint64) []byte {
	b := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	_ = b[7] //nice performance lol

	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)

	return b
}

func NoBoundsCheck(v uint64) []byte  {
	b := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)

	return b
}