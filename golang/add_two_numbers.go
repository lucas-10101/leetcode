package main

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode = &ListNode{}
	for _, e := range strconv.Itoa(nodeListToNumber(l1) + nodeListToNumber(l2)) {
		result.Val = int(e - 48)

		next := &ListNode{
			Val:  0,
			Next: result,
		}

		result = next
	}

	return result.Next
}

func nodeListToNumber(n *ListNode) int {

	result := 0
	nodeCopy := n
	mult := 1

	for nodeCopy != nil {
		result += nodeCopy.Val * mult
		mult *= 10
		nodeCopy = nodeCopy.Next
	}

	return result
}
