package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalize(t *testing.T) {
	type args struct {
		a        string
		b        string
		distance int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "zero length", args: args{a: "", b: "", distance: 10}, want: 10},
		{name: "ケーキ チョコレートケーキ", args: args{a: "ケーキ", b: "チョコレートケーキ", distance: 6}, want: float64(6) / float64(9)},
		{name: "abc def", args: args{a: "abc", b: "def", distance: 3}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Normalize(tt.args.a, tt.args.b, tt.args.distance), "Normalize(%v, %v, %v)", tt.args.a, tt.args.b, tt.args.distance)
		})
	}
}
