package Week_06

//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min2(dp[i-1][j], min2(dp[i][j-1], dp[i-1][j-1])) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}

		}
	}
	return maxSide * maxSide
}
func min2(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
