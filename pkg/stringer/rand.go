package stringer

import (
	"githu.com/weblfe/flyfire/pkg/random"
	"math"
	"math/rand"
	"sort"
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

type PwdLevel struct {
	Length  uint
	Number  [2]uint
	Special [2]uint
	Lower   [2]uint
	Upper   [2]uint
}

func NewPwdLevelDefault() PwdLevel {
	return PwdLevel{
		Length:  10,
		Number:  [2]uint{1, 3},
		Special: [2]uint{2, 3},
		Lower:   [2]uint{3, 5},
		Upper:   [2]uint{1, 3},
	}
}

func NewPwdLevel(length uint) PwdLevel {
	return PwdLevel{
		Length:  length,
		Number:  [2]uint{1, uint(math.Ceil(float64(length) * 1/10))},
		Special: [2]uint{1, uint(math.Ceil(float64(length) * 1/10))},
		Lower:   [2]uint{1, uint(math.Ceil(float64(length) * 2/5))},
		Upper:   [2]uint{1, uint(math.Ceil(float64(length) * 2/5))},
	}
}

func (l *PwdLevel) calculate() {
	var (
		left  = l.Length
		index []*[2]uint
	)
	if l.Number[1] > 0 {
		if l.Number[0] > left {
			l.Number[0] = random.Int(l.Number[0]-left, left)
		}
		if l.Number[1] > left {
			l.Number[1] = minUint(l.Number[1]-left, left)
		}
		index = append(index, &l.Number)
		left = left - l.Number[1]
	}
	if l.Special[1] > 0 {
		if l.Special[0] > left {
			l.Special[0] = random.Int(l.Special[0]-left, left)
		}
		if l.Special[1] > left {
			l.Special[1] = minUint(l.Special[1]-left, left)
		}
		index = append(index, &l.Special)
		left = left - l.Special[1]
	}
	if l.Lower[1] > 0 {
		if l.Lower[0] > left {
			l.Lower[0] = random.Int(l.Lower[0]-left, left)
		}
		if l.Lower[1] > left {
			l.Lower[1] = minUint(l.Lower[1]-left, left)
		}
		index = append(index, &l.Lower)
		left = left - l.Lower[1]
	}
	if l.Upper[1] > 0 {
		if l.Upper[0] > left {
			l.Upper[0] = random.Int(l.Upper[0]-left, left)
		}
		if l.Upper[1] > left {
			l.Upper[1] = left
		}
		index = append(index, &l.Upper)
		left = left - l.Upper[1]
	}
	sort.Slice(index, func(i, j int) bool {
		if (*index[i])[1] >= (*index[j])[1] {
			return true
		}
		return false
	})
	if left > 0 {
		(*index[0])[1] = index[0][1] + left
		return
	}
}

func GeneratePassword(cnf PwdLevel) string {
	var (
		number  = []rune("0123456789")
		special = []rune("-_#!@#%^&*()<>?.:;")
		lower   = []rune("abcdefghijklmnopqrstuvwxyz")
		upper   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	)
	var (
		rs      []rune
		handler = func(src []rune, end uint) []rune {
			var arr []rune
			size := len(src)
			// 不足补充
			if end > uint(size) {
				// @todo
			}
			rand.Shuffle(size, func(i, j int) {
				src[i], src[j] = src[j], src[i]
			})
			return append(arr, src[:end]...)
		}
	)
	// 参数调整计算
	cnf.calculate()
	if cnf.Number[1] > 0 {
		rs = append(rs, handler(number, cnf.Number[1])...)
	}
	if cnf.Special[1] > 0 {
		rs = append(rs, handler(special, cnf.Special[1])...)
	}
	if cnf.Lower[1] > 0 {
		rs = append(rs, handler(lower, cnf.Lower[1])...)
	}
	if cnf.Upper[1] > 0 {
		rs = append(rs, handler(upper, cnf.Upper[1])...)
	}
	rand.Shuffle(int(cnf.Length), func(i, j int) {
		rs[i], rs[j] = rs[j], rs[i]
	})
	return string(rs)
}

func minUint(v, v2 uint) uint {
	if v <= v2 {
		return v
	}
	return v2
}
