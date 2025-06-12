package main

import "fmt"

func main() {

	str := "cbbd"
	fmt.Println(longestPalindrome(str))
}

// https://leetcode.cn/problems/longest-palindromic-substring
// 最短回文，采取动态规划解题，f(i,j) = f(i+1, j-1) ^ Si == Sj。其中迭代回合计数，是当前字符串长度，从2开始，字符串长为1，默认就是回文
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var dp = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 单个字符永远是回文
	}

	var longestStr string = s[0:1]
	var longest int = 1

	// 迭代回合计数，是当前字符串长度
	for length := 2; length <= len(s); length++ {
		for i := 0; i < len(s); i++ {
			j := i + length - 1
			if j >= len(s) { // j越界，退出当前判定
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}

			if length <= 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] && (j-i+1) >= longest {
				longestStr = s[i : j+1]
			}
		}
	}

	return longestStr
}

func reverseStr(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(s[i])
	}
	return res
}

// 正向全遍历，暴力解法，超时
func longestPalindrome_v1(s string) string {
	if len(s) <= 1 {
		return s
	}

	var longestStr string = s[0:1]
	var longest int

	for i, _ := range s {
		for j := i + 1; 2*j-i <= len(s)+1; j++ {
			if (2*j-i <= len(s)) && s[i:j] == reverseStr(s[j:2*j-i]) && longest <= 2*(j-i) {
				longest = 2 * (j - i)
				longestStr = s[i : 2*j-i]
			}

			if (j != i+1) && s[i:j-1] == reverseStr(s[j:2*j-i-1]) && longest <= 2*(j-i)-1 {
				longest = 2*(j-i) - 1
				longestStr = s[i : 2*j-i-1]
			}
		}
	}

	return longestStr
}
