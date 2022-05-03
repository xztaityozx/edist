package lib

import "github.com/samber/lo"

type Levenshtein struct {
	ReplaceCost int
}

func NewLevenshtein(cost int) Levenshtein {
	return Levenshtein{ReplaceCost: cost}
}

func (l Levenshtein) Walk(a, b string) (int, error) {
	aRunes := []rune(a)
	bRunes := []rune(b)
	aLen := len(aRunes)
	bLen := len(bRunes)
	dp := createTable(aLen, bLen)

	for i := 1; i <= aLen; i++ {
		for k := 1; k <= bLen; k++ {
			cost := l.ReplaceCost
			if aRunes[i-1] == bRunes[k-1] {
				cost = 0
			}

			dp[i][k] = lo.Min([]int{
				dp[i-1][k] + 1,
				dp[i][k-1] + 1,
				dp[i-1][k-1] + cost,
			})
		}
	}

	return dp[aLen][bLen], nil
}
