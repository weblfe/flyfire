package random_test

import (
		"fmt"
		"githu.com/weblfe/flyfire/pkg/random"
		"github.com/stretchr/testify/assert"
		"math"
		"testing"
)

func TestInt(t *testing.T) {
		var (
				as = assert.New(t)
				cases = []struct{
						min  uint
						max  uint
				}{
						{
								min: 0,
								max: 10,
						},
						{
								min: 1,
								max: 20,
						},
						{
								min: 0,
								max: math.MaxInt,
						},
						{
								min: 10000000,
								max: 20,
						},
				}
		)
		for _,v:=range cases {
				t.Run(fmt.Sprintf("random.Int(%d,%d)",v.min,v.max), func(t *testing.T) {
						r:=random.Int(v.min,v.max)
						t.Logf("r=%d",r)
						as.True(r>=min(v.min,v.max) && r<=max(v.min,v.max),"随机数异常" )
				})
		}
}

func min(v,v2 uint) uint  {
		if v <= v2 {
				return v
		}
		return v2
}

func max(v,v2 uint) uint  {
		if v >= v2 {
				return v
		}
		return v2
}
