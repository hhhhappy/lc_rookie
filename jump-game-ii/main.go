package main

import "fmt"

func main() {
	res := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(res))
}

// https://leetcode.cn/problems/jump-game-ii
// 基本的DP思路，从前往后，遍历出所有idx的最小步数，然后最小值累加1。判定条件：nums[j]+j >= i，不可达以-1表示
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		reach := false
		min := len(nums)
		for j := 0; j < i; j++ {
			if dp[j] >= 0 && nums[j]+j >= i {
				reach = true
				if min >= dp[j]+1 {
					min = dp[j] + 1
				}
			}
		}
		if reach {
			dp[i] = min
		} else {
			dp[i] = -1
		}
	}
	return dp[len(dp)-1]
}
