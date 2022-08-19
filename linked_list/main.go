package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	val  int
	next *ListNode
}

func convertToLL(ll []int) *ListNode {
	var root, prev *ListNode
	for i := len(ll) - 1; i >= 0; i-- {
		root = &ListNode{
			val:  ll[i],
			next: prev,
		}

		prev = root
	}

	return root
}

func traverse(root *ListNode) {
	for ; root != nil; root = root.next {
		fmt.Println(root.val)
	}
}

func reverse(root *ListNode) *ListNode {
	var prev *ListNode
	var temp *ListNode

	for ; root != nil; root = temp {
		temp = root.next
		root.next = prev
		prev = root
	}

	return prev
}

// Calculate sum of 2 integers expressed through linked lists
func calcSum(l1 *ListNode, l2 *ListNode) int {
	l1 = reverse(l1)
	l2 = reverse(l2)

	sum := 0
	i := 0
	var digSum, carry int

	for ; l1 != nil && l2 != nil; l1, l2 = l1.next, l2.next {
		digSum = l1.val + l2.val + carry
		carry = digSum / 10
		digSum = digSum % 10
		sum += digSum * int(math.Pow10(i))
		i++
	}

	for ; l1 != nil; l1 = l1.next {
		digSum = l1.val + carry
		carry = digSum / 10
		digSum = digSum % 10
		sum += digSum * int(math.Pow10(i))
		i++
	}

	for ; l2 != nil; l2 = l2.next {
		digSum = l2.val + carry
		carry = digSum / 10
		digSum = digSum % 10
		sum += digSum * int(math.Pow10(i))
		i++
	}

	return sum + carry*int(math.Pow10(i))
}
