package stringer

import (
	"bytes"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Buffer struct {
	*bytes.Buffer
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.WriteString(strconv.Itoa(val))
	case int64:
		b.WriteString(strconv.FormatInt(val, 10))
	case uint:
		b.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.WriteString(strconv.FormatUint(val, 10))
	case string:
		b.WriteString(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// FilterBothSidesSpace 过滤两边的空格
func FilterBothSidesSpace(s string) string {
	spaces := []string{" ", "　", " "}
	for _, space := range spaces {
		s = strings.Trim(s, space)
	}
	return s
}

// WithNoNull 对比非空字符串返回
func WithNoNull(v, defaultVal string) string {
	if v == "" || strings.TrimSpace(v) == "" {
		return defaultVal
	}
	return v
}

// IsInt 是否整形数字字符串
func IsInt(v string) bool {
	if v == "" {
		return false
	}
	for _, char := range []rune(v) {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// IsNumeric 是否数字 eg: 1,0, 0.1, 00.11,1.00, 099,900
// @examples IsNumeric("1") == true
func IsNumeric(v string) bool {
	if strings.Count(v, ".") > 1 {
		return false
	}
	for _, char := range []rune(strings.ReplaceAll(v, ".", "")) {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func LT(a, b string) bool {
	return a < b
}

func GT(a, b string) bool {
	return a > b
}

type RandomArg struct {
	Length   int
	SpecChar []string
}

type Random struct {
	arg     *RandomArg
	handler map[string]func(length int, specChar ...string) string
}

func NewRandomArg(len int, specChar []string) *RandomArg {
	if len <= 0 {
		len = 8
	}
	return &RandomArg{
		Length:   len,
		SpecChar: specChar,
	}
}

func NewRandom(option ...func(arg *RandomArg)) *Random {
	var arg = NewRandomArg(8, nil)
	if len(option) > 0 {
		for _, o := range option {
			o(arg)
		}
	}
	return &Random{
		arg: arg,
		handler: map[string]func(length int, specChar ...string) string{
			"number": RandomNumber,
			"string": GenerateRandomStr,
		},
	}
}

func (r *Random) Random(name string, length ...int) string {
	handler, ok := r.handler[name]
	if !ok {
		return ``
	}
	if len(length) > 0 && length[0] > 0 {
		return handler(length[0], r.arg.SpecChar...)
	}
	return handler(r.arg.Length, r.arg.SpecChar...)
}

// RandomNumber 随机数数字
func RandomNumber(length int, specChar ...string) string {
	var str = []rune("0123456789")
	if len(specChar) > 0 {
		for _, char := range specChar {
			str = append(str, []rune(char)...)
		}
	}
	var (
		rs   []rune
		size = len(str)
	)
	// 随机打乱
	rand.Shuffle(size, func(i, j int) {
		str[i], str[j] = str[j], str[i]
	})
	var seed = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		rs = append(rs, str[seed.Intn(size)])
	}
	return string(rs)
}

// SplitLast 分隔字符串获取最后一个
func SplitLast(value string, sep ...string) string {
	sep = append(sep, ",")
	if !strings.Contains(value, sep[0]) {
		return value
	}
	values := strings.Split(value, sep[0])
	return values[len(values)-1]
}

// SplitFirst 分隔字符串获取最后一个
func SplitFirst(value string, sep ...string) string {
	sep = append(sep, ",")
	if !strings.Contains(value, sep[0]) {
		return value
	}
	values := strings.Split(value, sep[0])
	return values[0]
}

func Contains(ss []string, sub string) bool {
	for _, s := range ss {
		if s == sub {
			return true
		}
	}
	return false
}
