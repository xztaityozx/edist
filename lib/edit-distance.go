package lib

type EditDistance interface {
  Walk(a,b string) (int, error)
}

func createTable(a, b int) [][]int {
  dp := make([][]int,a+1)
  for i:=0;i<a+1;i++ {
    dp[i] = make([]int, b+1)
  }

  for i:=0;i<a+1;i++ {
    dp[i][0] = i
  }
  for i:=0;i<b+1;i++ {
    dp[0][i] = i
  }

  return dp
}
