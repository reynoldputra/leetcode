package main

import "fmt"

const M = 1000000007

func rec(total int32, k int32, dp [][]int32) int32 {
  if dp[total][k] != 0 {
    return dp[total][k]
  }

  if total == 0 || k == 1 {
    return 1
  }

  var c int32 = 0
  for j := int32(1); j <= k; j ++ {
    if total >= j {
      r := rec(total - j, j, dp) 
      c = (c + r) % M
    }
  }

  return c
}


func cdp(total int32, k int32, dp [][]int32) {
  for i := int32(1); i <= total; i ++ {
    for j := int32(1); j <= k; j ++ {
      if i >= j {
        dp[i][j] = rec(i, j, dp) % M
      }
    }
  }
}

func main () {
  total := int32(842)
  k := int32(91)
  dp := make([][]int32, total+1)
	for i := range dp {
		dp[i] = make([]int32, k+1)
	}

  cdp(total, k, dp)
  fmt.Println(dp[total])

  fmt.Println(dp[total][k])
}
