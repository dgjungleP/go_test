package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	} else {
		node = nil
	}
}

func main() {

}
