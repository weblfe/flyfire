package stringer_test

import (
		"bytes"
		"fmt"
		"github.com/stretchr/testify/assert"
		"github.com/weblfe/flyfire/pkg/stringer"
		"reflect"
		"testing"
)

func TestBuffer_Append(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		i interface{}
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   *stringer.Buffer
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &stringer.Buffer{
				Buffer: tt.fields.Buffer,
			}
			if got := b.Append(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamel2Case(t *testing.T) {
	type args struct {
		name string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringer.Camel2Case(tt.args.name); got != tt.want {
				t.Errorf("Camel2Case() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterBothSidesSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "　name ",
			},
			want: "name",
		},
		{
			args: args{s: "special space "},
			want: "special space",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringer.FilterBothSidesSpace(tt.args.s); got != tt.want {
				t.Errorf("FilterBothSidesSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomStr(t *testing.T) {
	type args struct {
		length int
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringer.GenerateRandomStr(tt.args.length); got != tt.want {
				t.Errorf("GenerateRandomStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBuffer(t *testing.T) {
	var tests []struct {
		name string
		want *stringer.Buffer
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringer.NewBuffer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	var (
		as    = assert.New(t)
		cases = []struct {
			level stringer.PwdLevel
		}{
			{
				level: stringer.NewPwdLevelDefault(),
			},
			{
				level: stringer.NewPwdLevel(6),
			},
			{
				level: stringer.NewPwdLevel(9),
			},
			{
				level: stringer.NewPwdLevel(8),
			},
			{
				level: stringer.NewPwdLevel(11),
			},
			{
				level: stringer.NewPwdLevel(64),
			},
		}
	)
	for _, v := range cases {
		t.Run(fmt.Sprintf("%v", v.level), func(t *testing.T) {
			pwd := stringer.GeneratePassword(v.level)
			t.Logf("pwd=%v", pwd)
			as.Equal(v.level.Length, uint(len(pwd)), "生成密码长度不正确")
		})
	}
}

func BenchmarkGeneratePassword(b *testing.B) {
	var (
		as    = assert.New(b)
		cases = []struct {
			level stringer.PwdLevel
		}{
			{
				level: stringer.NewPwdLevelDefault(),
			},
		}
	)
	b.ResetTimer()
	b.N = 10000
	for _, v := range cases {
		b.Run(fmt.Sprintf("%v", v.level), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				pwd := stringer.GeneratePassword(v.level)
				t.Logf("pwd=%v", pwd)
				as.Equal(v.level.Length, uint(len(pwd)), "生成密码长度不正确")
			}
		})
	}
}
