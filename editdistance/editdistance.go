package editdistance

// EditDistance O(|a|*|b|)
func EditDistance(a, b string) int64 {
	n := len(a) + 1
	m := len(b) + 1
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][0] = 0

	min := func(x, y int64) int64 {
		if x < y {
			return x
		}
		return y
	}

	for i, x := range a {
		for j, y := range b {
			if x == y {
				dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j])
			} else {
				dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j]+1)
			}
			dp[i][j+1] = min(dp[i][j+1], dp[i][j]+1)
			dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
		}
	}

	return dp[len(a)][len(b)]
}

// LevenshteinDistance is alias of EditDistance
func LevenshteinDistance(a, b string) int64 {
	return EditDistance(a, b)
}
