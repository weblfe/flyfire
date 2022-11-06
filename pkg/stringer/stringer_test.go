package stringer

import (
	"bytes"
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
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Buffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
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
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel2Case(tt.args.name); got != tt.want {
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
			if got := FilterBothSidesSpace(tt.args.s); got != tt.want {
				t.Errorf("FilterBothSidesSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomStr(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomStr(tt.args.length); got != tt.want {
				t.Errorf("GenerateRandomStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBuffer(t *testing.T) {
	tests := []struct {
		name string
		want *Buffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuffer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
