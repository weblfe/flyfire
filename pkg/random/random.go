package random

import (
		"crypto/rand"
		"math"
		"math/big"
)

func Int(min, max uint) uint {
	if min > max {
			min,max = max,min
	}
	v:= max-min
	if v < math.MaxInt64 {
			v = v+ 1
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(v)))
	num := uint(result.Int64())
	return num + min
}
