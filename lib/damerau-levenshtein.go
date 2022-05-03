package lib

import "github.com/samber/lo"

type DamerauLevenshtein struct {
  Levenshtein
  SwapCost int
  Limited bool
}

func NewDamerauLevenshtein(replaceCost, swapCost int, limited bool) DamerauLevenshtein {
  return DamerauLevenshtein{Limited: limited, SwapCost: swapCost, Levenshtein: NewLevenshtein(replaceCost)}
}

func (d DamerauLevenshtein) limitedWalk(a, b string) (int, error) {
  aLen := len(a)
  bLen := len(b)
  dp := createTable(aLen, bLen)

  for i:=1; i <= aLen; i++ {
    for k:=1;k<=bLen;k++ {
      cost := d.Levenshtein.ReplaceCost
      if a[i-1] == b[k-1] {
        cost = 0
      }

      dp[i][k] = lo.Min([]int{
        dp[i-1][k  ] + 1,
        dp[i  ][k-1] + 1,
        dp[i-1][k-1] + cost,
      })

      if i > 1 && k > 1 && a[i-1] == b[k-2] && a[i-2] == b[k-1] {
        dp[i][k] = lo.Min([]int{dp[i][k], dp[i-2][k-2]+d.SwapCost})
      }
    }
  }

  return dp[aLen][bLen], nil
}

func (d DamerauLevenshtein) walk(a, b string) (int, error) {
  aLen, bLen := len(a), len(b)
  max := aLen+bLen

  dict := map[rune]int{}

  dp := map[int]map[int]int{}
  for i := -1; i<aLen+1;i++ {
    dp[i] = map[int]int{}
  }

  dp[-1][-1] = max

  for i:=0;i<aLen+1;i++ {
    dp[i][-1] = max
    dp[i][0] = i
  }
  for i:=0;i<bLen+1;i++ {
    dp[-1][i] = max
    dp[0][i] = i
  }

  for i:=1; i<aLen+1; i++ {
    db := 0
    for j:=1; j<bLen+1; j++ {
      k := dict[rune(b[j-1])]
      l := db

      cost := d.ReplaceCost
      if a[i-1] == b[j-1] {
        db = j
        cost = 0
      }

      dp[i][j] = lo.Min([]int{
        dp[i-1][j-1] + cost,
        dp[i  ][j-1] + 1,
        dp[i-1][j  ] + 1,
        dp[k-1][l-1] + (i-k-1) + d.SwapCost + (j-l-1),
      })
    }
    dict[rune(a[i-1])] = i
  }

  return dp[aLen][bLen], nil
}

func (d DamerauLevenshtein) Walk(a, b string) (int, error) {
  if d.Limited {
    return d.limitedWalk(a, b)
  } else {
    return d.walk(a, b)
  }
}
