package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDamerauLevenshtein(t *testing.T) {
	type args struct {
		replaceCost int
		swapCost    int
		isLimited   bool
	}

	tests := []struct {
		name   string
		args   args
		expect DamerauLevenshtein
	}{
		{name: "replaceCost 10, limited", args: args{replaceCost: 10, swapCost: 11, isLimited: true}, expect: DamerauLevenshtein{Levenshtein: Levenshtein{ReplaceCost: 10}, SwapCost: 11, Limited: true}},
		{name: "replaceCost 10, unlimited", args: args{replaceCost: 10, swapCost: 11, isLimited: false}, expect: DamerauLevenshtein{Levenshtein: Levenshtein{ReplaceCost: 10}, SwapCost: 11, Limited: false}},
	}

	for _, tt := range tests {
		actual := NewDamerauLevenshtein(tt.args.replaceCost, tt.args.swapCost, tt.args.isLimited)
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestDamerauLevenshtein_Walk(t *testing.T) {
	type args struct {
		replaceCost int
		swapCost    int
		limited     bool
		a           string
		b           string
	}

	tests := []struct {
		args     args
		expect   int
		hasError bool
	}{
		{args: args{replaceCost: 1, swapCost: 1, limited: true, a: "str", b: "stu"}, expect: 1, hasError: false},
		{args: args{replaceCost: 1, swapCost: 1, limited: true, a: "ca", b: "abc"}, expect: 3, hasError: false},
		{args: args{replaceCost: 1, swapCost: 2, limited: true, a: "str", b: "stu"}, expect: 1, hasError: false},
		{args: args{replaceCost: 1, swapCost: 2, limited: true, a: "ca", b: "abc"}, expect: 3, hasError: false},
		{args: args{replaceCost: 1, swapCost: 1, limited: true, a: "acb", b: "abc"}, expect: 1, hasError: false},
		{args: args{replaceCost: 1, swapCost: 3, limited: true, a: "acb", b: "abc"}, expect: 2, hasError: false},
		{args: args{replaceCost: 1, swapCost: 3, limited: true, a: "acb", b: "abc"}, expect: 2, hasError: false},
		{args: args{replaceCost: 1, swapCost: 3, limited: true, a: "キュアミラクル", b: "キュアマジカル"}, expect: 3, hasError: false},

		{args: args{replaceCost: 1, swapCost: 1, limited: false, a: "str", b: "stu"}, expect: 1, hasError: false},
		{args: args{replaceCost: 1, swapCost: 2, limited: false, a: "str", b: "stu"}, expect: 1, hasError: false},
		{args: args{replaceCost: 1, swapCost: 1, limited: false, a: "ca", b: "abc"}, expect: 2, hasError: false},
		{args: args{replaceCost: 1, swapCost: 1, limited: false, a: "acb", b: "abc"}, expect: 1, hasError: false},
		{args: args{replaceCost: 1, swapCost: 3, limited: false, a: "acb", b: "abc"}, expect: 2, hasError: false},
		{args: args{replaceCost: 1, swapCost: 3, limited: false, a: "キュアミラクル", b: "キュアマジカル"}, expect: 3, hasError: false},
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("%s <-> %s cost=(%v,%v) limited: %v", v.args.a, v.args.b, v.args.replaceCost, v.args.swapCost, v.args.limited), func(t *testing.T) {
			actual, err := NewDamerauLevenshtein(v.args.replaceCost, v.args.swapCost, v.args.limited).Walk(v.args.a, v.args.b)
			if (err != nil) != v.hasError {
				assert.Errorf(t, err, "DamerauLevenshtein.Walk() hasError = %v", v.hasError)
				return
			}
			assert.Equal(t, v.expect, actual)
		})
	}
}
