package lib

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLevenshtein(t *testing.T) {
	cost := 10
	expect := Levenshtein{ReplaceCost: cost}
	actual := NewLevenshtein(cost)
	assert.Equal(t, expect, actual)
}

func TestLevenshtein_Walk(t *testing.T) {
	type arg struct {
		a string
		b string
	}
	tests := []struct {
		name     string
		arg      arg
		expect   int
		cost     int
		hasError bool
	}{
		{name: "str <-> stu", arg: arg{a: "str", b: "stu"}, expect: 1, cost: 1, hasError: false},
		{name: "str <-> stir", arg: arg{a: "str", b: "stir"}, expect: 1, cost: 1, hasError: false},
		{name: "kitten <-> sitting", arg: arg{a: "kitten", b: "sitting"}, expect: 3, cost: 1, hasError: false},
		{name: "str <-> abc", arg: arg{a: "str", b: "abc"}, expect: 3, cost: 1, hasError: false},
		{name: "str <-> str", arg: arg{a: "str", b: "str"}, expect: 0, cost: 1, hasError: false},
		{name: "str <-> ", arg: arg{a: "str", b: ""}, expect: 3, cost: 1, hasError: false},

		{name: "str <-> stu", arg: arg{a: "str", b: "stu"}, expect: 2, cost: 2, hasError: false},
		{name: "str <-> stir", arg: arg{a: "str", b: "stir"}, expect: 1, cost: 2, hasError: false},
		{name: "kitten <-> sitting", arg: arg{a: "kitten", b: "sitting"}, expect: 5, cost: 2, hasError: false},
		{name: "str <-> abc", arg: arg{a: "str", b: "abc"}, expect: 6, cost: 2, hasError: false},
		{name: "str <-> str", arg: arg{a: "str", b: "str"}, expect: 0, cost: 2, hasError: false},
		{name: "str <-> ", arg: arg{a: "str", b: ""}, expect: 3, cost: 2, hasError: false},
		{name: "キュアミラクル <-> キュアマジカル", arg: arg{a: "キュアミラクル", b: "キュアマジカル"}, expect: 3, cost: 1, hasError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Levenshtein{ReplaceCost: tt.cost}.Walk(tt.arg.a, tt.arg.b)
			if (err != nil) != tt.hasError {
				t.Errorf("Levenshtein.Walk() want error = %v, hasError = %v", err, tt.hasError)
				return
			}

			if !reflect.DeepEqual(actual, tt.expect) {
				t.Errorf("Levenshtein.Walk() actual = %v expect = %v", actual, tt.expect)
			}
		})
	}
}
