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
  var dp = make([][]int, aLen+1)
  for i := 0;i<aLen+1;i++ {
    dp[i] = make([]int, bLen+1)
  }

  for i:=0;i<aLen+1;i++ {
    dp[i][0] = i;
  }
  for i:=0;i<bLen+1;i++ {
    dp[0][i] = i;
  }

  for i := 1; i<= aLen; i++ {
    for k := 1; k<= bLen; k++ {
      cost := l.ReplaceCost;
      if a[i-1] == b[k-1] {
        cost = 0
      }

      dp[i][k] = lo.Min([]int{
        dp[i-1][k  ] + 1, 
        dp[i  ][k-1] + 1,
        dp[i-1][k-1] + cost,
      })
    }
  }

  return dp[aLen][bLen], nil
}
