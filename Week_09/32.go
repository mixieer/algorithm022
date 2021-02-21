package Week_09

func longestValidParentheses(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	str := []byte(s)
	dp := []int{}
	for i := 0; i < len(str); {
		dp[i] = 0
		if str[i] == ')' {
			continue
		}
		j := i + 1
		if j < len(str) && str[i+1] == ')' {
			dp[i] = max(dp[i-2]+1, dp[i])
		}
		i = i + 2
	}
	var maxDp int
	for v := range dp {
		if v > maxDp {
			maxDp = v
		}
	}
	return maxDp
}
