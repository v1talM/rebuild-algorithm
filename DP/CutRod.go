package DP

var p = []int{1,5,8,9,10,17,17,20,24,30}
func CutRod(n int) int {
	lens := len(p)
	var store = []int{}
	for i := 0; i <= lens; i++ {
		store = append(store, -0x7fffffff)
	}
	return cutRod(store, n)
}

func cutRod(store []int, n int) int {
	if store[n] >= 0 {
		return store[n]
	}
	var q int
	if n == 0 {
		store[n] = 0
	} else {
		q = -0x7fffffff
		for i := 0; i < n; i++ {
			q = Max(q, p[i] + cutRod(store, n - i - 1))
		}
	}
	store[n] = q
	return q
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}