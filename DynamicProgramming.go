package main

import (
	"fmt"
	"rebuild-algorithm/DP"
)

func main() {
	max := DP.CutRod(4)
	fmt.Println(max)
	trangle := DP.MaxLen()
	fmt.Println(trangle)
	lcs := DP.LCS("-ABDFEG", "-AFDG")
	fmt.Println(lcs)
}

