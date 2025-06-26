package main

import "fmt"

// https://leetcode.cn/problems/surrounded-regions/submissions/639334915/?envType=problem-list-v2&envId=breadth-first-search
// 广度优先， 利用队列先入先出，去遍历一遍好节点，然后好节点置为x
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	var badNodes = map[string][]int{}
	for i := 0; i < m; i++ {
	L1:
		for j := 0; j < n; j++ {
			if _, ok := badNodes[genXY(i, j)]; !ok && board[i][j] == 'O' {
				var island = map[string][]int{
					genXY(i, j): {i, j},
				}
				var queue = [][]int{
					{i, j},
				}

				countIsland := func(x, y int) {
					if _, ok := island[genXY(x, y)]; !ok && x < m && x >= 0 && y < n && y >= 0 && board[x][y] == 'O' {
						island[genXY(x, y)] = []int{x, y}
						queue = append(queue, []int{x, y})
					}
				}

				for len(queue) != 0 {
					curx, cury := queue[0][0], queue[0][1]
					if !isGoodNode(board, curx, cury) {
						// 路径中存在异常node，跳过
						for key, _ := range island {
							badNodes[key] = island[key]
						}
						continue L1
					}
					countIsland(curx+1, cury)
					countIsland(curx-1, cury)
					countIsland(curx, cury+1)
					countIsland(curx, cury-1)
					queue = queue[1:]
				}

				for _, node := range island {
					board[node[0]][node[1]] = 'X'
				}
			}
		}
	}
}

func genXY(x, y int) string {
	return fmt.Sprintf("x%dy%d", x, y)
}

func isGoodNode(board [][]byte, x, y int) bool {
	m := len(board)
	n := len(board[0])
	if board[x][y] != 'O' {
		return false
	}

	if x+1 >= m || y+1 >= n || y-1 < 0 || x-1 < 0 {
		// 节点靠边
		return false
	}

	var up, down, left, right bool
	if board[x+1][y] == 'X' || board[x+1][y] == 'O' {
		up = true
	}

	if board[x-1][y] == 'X' || board[x-1][y] == 'O' {
		down = true
	}

	if board[x][y+1] == 'X' || board[x][y+1] == 'O' {
		left = true
	}
	if board[x][y-1] == 'X' || board[x][y-1] == 'O' {
		right = true
	}

	return up && down && left && right
}
