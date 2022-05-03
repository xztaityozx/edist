package lib

import (
	"errors"
	"github.com/samber/lo"
)

type HammingDistance struct{}

func NewHammingDistance() HammingDistance {
	return HammingDistance{}
}

func (h HammingDistance) Walk(a, b string) (int, error) {
	aRunes, bRunes := []rune(a), []rune(b)
	if len(aRunes) != len(bRunes) {
		return -1, errors.New("two string must be same length")
	}

	cnt := 0
	for _, v := range lo.Zip2(aRunes, bRunes) {
		if v.A != v.B {
			cnt++
		}
	}

	return cnt, nil
}
