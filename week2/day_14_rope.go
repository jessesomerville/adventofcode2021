package week2

type LinkedList struct {
	Head *LLNode
}

func FromString(s string) *LinkedList {
	ll := &LinkedList{
		Head: &LLNode{Val: rune(s[0])},
	}
	prevNode := ll.Head
	for _, r := range s[1:] {
		prevNode.InsertAfter(r)
		prevNode = prevNode.Next
	}
	return ll
}

type LLNode struct {
	Val  rune
	Next *LLNode
}

func (n *LLNode) InsertAfter(r rune) {
	newNode := &LLNode{
		Val:  r,
		Next: n.Next,
	}
	n.Next = newNode
}

func (ll *LinkedList) String() string {
	currNode := ll.Head
	str := ""
	for currNode != nil {
		str += string(currNode.Val)
		currNode = currNode.Next
	}
	return str
}
