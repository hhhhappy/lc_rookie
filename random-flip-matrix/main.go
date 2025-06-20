package main

import "math/rand"

func main() {

}

// https://leetcode.cn/problems/random-flip-matrix/description/?envType=problem-list-v2&envId=hash-table
// 转一维数组，total = n * m - 1，total对于n取余则是y轴，取整则是X轴
// 每当某一个数字被随机出队，那么就在store中存储其映射(n， total)。若store存在total的映射那么就存储(n, store(total))
type Solution struct {
	m, n, total int
	store       map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{
		store: map[int]int{},
		m:     m,
		n:     n,
		total: m*n - 1,
	}
}

func (this *Solution) Flip() []int {
	randV := rand.Intn(this.total + 1)
	var index int
	if v, ok := this.store[randV]; ok {
		// 已经被替换过，返回替换坐标
		index = v
	} else {
		index = randV
	}

	if v, ok := this.store[this.total]; ok {
		this.store[randV] = v
	} else {
		this.store[randV] = this.total
	}

	this.total--
	return []int{index / this.n, index % this.n}
}

func (this *Solution) Reset() {
	this.store = map[int]int{}
	this.total = this.m*this.n - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
