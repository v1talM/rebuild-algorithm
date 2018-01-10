package DP

var maxLen = [7][6]int{}
func LCS(s1, s2 string) int {
	for i := 0; i < len(s1); i++ {
		maxLen[i][0] = 0
	}
	for j := 0; j < len(s2); j++ {
		maxLen[0][j] = 0
	}
	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			if s1[i - 1] == s2[j - 1] {
				maxLen[i][j] = maxLen[i - 1][j - 1] + 1
			} else {
				maxLen[i][j] = Max(maxLen[i - 1][j], maxLen[i][j - 1])
			}
		}
	}
	return maxLen[len(s1) - 1][len(s2) - 1]
}