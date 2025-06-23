package main

import (
	"math"
	"sort"
)

func main() {

}

// 贪心，但实际上是考察中位数和回文的判定。中位数最近的回文的上界和下界，并不能有效甄别更优解，所以需要重复计算，比大小
func minimumCost(nums []int) int64 {
	sort.Ints(nums)
	var a float64
	if len(nums)%2 != 0 {
		a = float64(nums[len(nums)/2])
	} else {
		a = (float64(nums[len(nums)/2-1]) + float64(nums[len(nums)/2])) / 2
	}

	lower, upper := nearestRightVal(a)
	var costU int64 = 0
	for i := 0; i < len(nums); i++ {
		costU += int64(math.Abs(float64(nums[i] - upper)))
	}

	var costL int64 = 0
	for i := 0; i < len(nums); i++ {
		costL += int64(math.Abs(float64(nums[i] - lower)))
	}
	if costL < costU {
		return costL
	}

	return costU
}

func isRightVal(num int) bool {
	if num < 0 {
		return false
	}
	vals := []int{}
	for {
		base := 10
		vals = append(vals, num%base)
		num = num / base
		if num == 0 {
			break
		}
	}

	i := 0
	j := len(vals) - 1
	for {
		if vals[i] != vals[j] {
			return false
		}

		i++
		j--

		if j <= i {
			break
		}
	}

	return true
}

func nearestRightVal(num float64) (int, int) {
	floor := math.Floor(num)
	if num-floor > 0.5 {
		floor += 1
	}

	target := int(floor)
	upper := -1
	lower := -1
	for i := 0; ; i++ {
		if upper < 0 && isRightVal(target+i) {
			upper = target + i
		}

		if lower < 0 && isRightVal(target-i) {
			lower = target - i
		}

		if upper >= 0 && lower >= 0 {
			return lower, upper
		}

	}
}
