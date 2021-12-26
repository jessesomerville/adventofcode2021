package week4

const (
	Empty = iota
	Right
	Down
)

type LinkedList struct {
	Type int
	Next *LinkedList
}
