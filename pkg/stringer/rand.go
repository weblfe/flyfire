package stringer

import (
	"math/rand"
	"time"
)

func GenerateRandomStr(length int, specChar ...string) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if len(specChar) != 0 && specChar[0] != "" {
		str += specChar[0]
	}
	bs := []byte(str)
	rs := make([]byte, 0, 64)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		rs = append(rs, bs[r.Intn(len(bs))])
	}

	return string(rs)
}

type PwdCfg struct {
	Length  uint
	Number  [2]uint
	Special [2]uint
	Lower   [2]uint
	Upper   [2]uint
}

func GeneratePassword(cnf PwdCfg) string {
	var (
		number  = []rune("0123456789")
		special = []rune("-_#!@#%^&*()<>?.:;")
		lower   = []rune("abcdefghijklmnopqrstuvwxyz")
		upper   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	)

	var rs []rune

	if cnf.Number[1] > 0 {
		rand.Shuffle(len(number), func(i, j int) {
			number[i], number[j] = number[j], number[i]
		})

	}
	if cnf.Special[1] > 0 {
		rand.Shuffle(len(special), func(i, j int) {
			special[i], special[j] = special[j], special[i]
		})
	}
	if cnf.Lower[1] > 0 {
		rand.Shuffle(len(lower), func(i, j int) {
			lower[i], lower[j] = lower[j], lower[i]
		})
	}
	if cnf.Upper[1] > 0 {
		rand.Shuffle(len(upper), func(i, j int) {
			upper[i], upper[j] = upper[j], upper[i]
		})
	}

	return string(rs)
}
