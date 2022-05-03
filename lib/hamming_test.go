package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistance_Walk(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{name: "zero length", args: args{a: "", b: ""}, want: 0, wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
			return false
		}},
		{name: "abc <-> def", args: args{a: "abc", b: "def"}, want: 3, wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
			return false
		}},
		{name: "1010 <-> 1011", args: args{a: "1010", b: "1011"}, want: 1, wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
			return false
		}},
		{name: "different length", args: args{a: "abc", b: ""}, want: -1, wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
			return err != nil
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HammingDistance{}
			got, err := h.Walk(tt.args.a, tt.args.b)
			if !tt.wantErr(t, err, fmt.Sprintf("Walk(%v, %v)", tt.args.a, tt.args.b)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Walk(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestNewHammingDistance(t *testing.T) {
	tests := []struct {
		name string
		want HammingDistance
	}{
		{name: "new", want: HammingDistance{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewHammingDistance(), "NewHammingDistance()")
		})
	}
}
