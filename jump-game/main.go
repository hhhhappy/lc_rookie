package main

func main() {

}

// https://leetcode.cn/problems/jump-game
// 基本的递归逻辑，方程f(n) = f(m+nums[m]), n为最后一个下标，m为上一跳
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	curIndex := len(nums) - 1
	for i := curIndex - 1; i >= 0; i-- {
		if i+nums[i] >= curIndex {
			return canJump(nums[:i])
		}
	}

	return false
}
