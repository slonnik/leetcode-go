package problem234

// https://leetcode.com/problems/palindrome-linked-list/
import (
	"github.com/slonnik/leetcode-go/util"
)

func isPalindrome(head *util.ListNode) bool {

	var invertedNode *util.ListNode
	node := head
	for node != nil {
		invertedNode = &util.ListNode{Val: node.Val, Next: invertedNode}
		node = node.Next
	}

	for head != nil {
		if head.Val != invertedNode.Val {
			return false
		}
		head = head.Next
		invertedNode = invertedNode.Next
	}

	return true
}
