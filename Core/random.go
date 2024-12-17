package core

import "math/rand"

func Random(min, max int) int {
	if max == 0 {
		return 0
	}
	return rand.Intn(max-min) + min
}
