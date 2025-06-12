package main

import "fmt"

func main() {

	fmt.Println(generate(4))
}

// https://leetcode.cn/problems/pascals-triangle
// 最基本的DP，其实就是誊写一遍逻辑，没有特别的
func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}
