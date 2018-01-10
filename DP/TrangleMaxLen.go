package DP

const n  = 6
var trangle = [n][n]int{{},{0,7},{0,3,8},{0,8,1,0},{0,2,7,4,4},{0,4,5,2,6,5}}
var store = [n][n]int{}

func MaxLen() int {
	for i := 0; i < n; i++ {
		store[n - 1][i] = trangle[n - 1][i]
	}
	for i := n - 2; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			store[i][j] = Max(store[i + 1][j], store[i + 1][j + 1]) + trangle[i][j]
		}
	}
	return store[1][1]
}
