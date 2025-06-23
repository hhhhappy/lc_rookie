package main

import "sort"

type person struct {
	costs []int
	val   int
}

// 贪心，但其实感觉和贪心没关系，就是将cost做差值，然后排序，前n个去A，后n个去B
//https://leetcode.cn/problems/two-city-scheduling/description
func twoCitySchedCost(costs [][]int) int {
	persons := make([]person, len(costs))
	for i, c := range costs {
		persons[i] = person{
			costs: costs[i],
			val:   c[0] - c[1],
		}
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].val <= persons[j].val
	})

	var res int
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			res += persons[i].costs[0]
		} else {
			res += persons[i].costs[1]
		}
	}
	return res
}
