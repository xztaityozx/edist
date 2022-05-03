package lib

import "github.com/samber/lo"

type EditDistance interface {
	Walk(a, b string) (int, error)
}

func createTable(a, b int) [][]int {
	dp := make([][]int, a+1)
	for i := 0; i < a+1; i++ {
		dp[i] = make([]int, b+1)
	}

	for i := 0; i < a+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < b+1; i++ {
		dp[0][i] = i
	}

	return dp
}

func Normalize(a, b string, distance int) float64 {
	aLen, bLen := len([]rune(a)), len([]rune(b))
	if aLen == 0 && bLen == 0 {
		return float64(distance)
	}

	max := lo.Max([]int{aLen, bLen})
	return float64(distance) / float64(max)
}
