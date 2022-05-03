package lib

import "github.com/samber/lo"

type Levenshtein struct {
	ReplaceCost int
}

func NewLevenshtein(cost int) Levenshtein {
	return Levenshtein{ReplaceCost: cost}
}

func (l Levenshtein) Walk(a, b string) (int, error) {
	aLen := len(a)
	bLen := len(b)
	dp := createTable(aLen, bLen)

	for i := 1; i <= aLen; i++ {
		for k := 1; k <= bLen; k++ {
			cost := l.ReplaceCost
			if a[i-1] == b[k-1] {
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
