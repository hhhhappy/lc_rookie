package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 栈，先进后出；更优解，利用双指针，head和tail相互距离N
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	nodeArr := []*ListNode{}

	for {
		nodeArr = append(nodeArr, cur)

		if cur.Next == nil {
			break
		}

		cur = cur.Next
	}

	if len(nodeArr) <= 1 {
		return nil
	}

	if len(nodeArr)-n <= 0 {
		return nodeArr[1]
	}

	if n == 1 {
		nodeArr[len(nodeArr)-2].Next = nil
	} else {
		nodeArr[len(nodeArr)-n-1].Next = nodeArr[len(nodeArr)-n+1]
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
