package random

import (
		"crypto/rand"
		"math/big"
)

func Int(min, max int) int {
	if min > max {
			min,max = max,min
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min) + 1))
	num := int(result.Int64())
	return num + min
}
