package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/?envType=problem-list-v2&envId=depth-first-search
func flatten(root *TreeNode) {
	subFlattenV2(root)
}

// 二叉树先序遍历，无额外空间，递归解法。需要利用subFlatten返回最后一个叶子节点，用于作为右子树的展开结果的父节点
func subFlatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := root.Right
	root.Right = root.Left
	root.Left = nil
	var lastLeaf *TreeNode
	if root.Right != nil {
		lastLeaf = subFlatten(root.Right)
	} else {
		lastLeaf = root
	}

	if right != nil {
		tmp := subFlatten(right)
		lastLeaf.Right = right
		lastLeaf = tmp
	}

	return lastLeaf
}

// 二叉树先序遍历，利用栈（先入后出）迭代解法。
func subFlattenV2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	stack := []*TreeNode{}
	list := []*TreeNode{}
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			list = append(list, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	for i, _ := range list {
		if i-1 >= 0 {
			list[i-1].Right = list[i]
			list[i-1].Left = nil
		}
	}

	return list[0]
}
