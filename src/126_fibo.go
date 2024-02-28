package main

import "fmt"

func Fibo(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	const mod = 1000000007
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
		fmt.Printf("dp[%d]:%d\n", i, dp[i])
	}
	return dp[n]
}
