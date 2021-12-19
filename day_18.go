package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	// RootVal is the value to give non-leaf nodes.
	RootVal = -1
)

var (
	//go:embed inputs/day_18.txt
	snailfishFile string
)

func snailfish() int {
	nums := strings.Split(snailfishFile, "\n")
	currNum := NewSnailfishNum(nums[0])
	for _, num := range nums[1:] {
		toAdd := NewSnailfishNum(num)
		added := currNum.Add(toAdd)
		added.Reduce()
		currNum = added
	}
	return currNum.Magnitude()
}

func snailfishLargest() int {
	numStrs := strings.Split(snailfishFile, "\n")
	nums := make([]*SnailfishNum, 0, len(numStrs))
	for _, n := range numStrs {
		nums = append(nums, NewSnailfishNum(n))
	}

	largestMag := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			a := nums[i].Copy()
			b := nums[j].Copy()

			added := a.Add(b)
			added.Reduce()
			mag := added.Magnitude()
			if mag > largestMag {
				largestMag = mag
			}
		}
	}
	return largestMag
}

// Copy returns a copy of the receiver.
func (sn SnailfishNum) Copy() *SnailfishNum {
	str := sn.String()
	return NewSnailfishNum(str)
}

// SnailfishNum represents a snailfish number pair.
type SnailfishNum struct {
	Value             int
	LeftNum, RightNum *SnailfishNum
	Parent            *SnailfishNum
	IsLeaf            bool
}

// Magnitude calculates the magnitude of the number.
func (sn *SnailfishNum) Magnitude() int {
	if sn.IsLeaf {
		return sn.Value
	}
	leftMag := sn.LeftNum.Magnitude()
	rightMag := sn.RightNum.Magnitude()
	return (3 * leftMag) + (2 * rightMag)
}

// Add returns the result of adding toAdd to the receiver.
func (sn *SnailfishNum) Add(toAdd *SnailfishNum) *SnailfishNum {
	added := &SnailfishNum{
		Value:    RootVal,
		LeftNum:  sn,
		RightNum: toAdd,
	}
	sn.Parent = added
	toAdd.Parent = added
	return added
}

// Reduce performs a reduction on the receiver.
func (sn *SnailfishNum) Reduce() {
	sn.ExplodeAll(sn, 0)
	for sn.SplitOne(sn) {
		sn.ExplodeAll(sn, 0)
	}
}

// ExplodeAll does every explode operation in the number.
func (sn *SnailfishNum) ExplodeAll(root *SnailfishNum, depth int) {
	if root == nil || root.IsLeaf {
		return
	}
	if depth >= 4 {
		if err := root.Explode(); err != nil {
			log.Fatal(err)
		}
	}
	sn.ExplodeAll(root.LeftNum, depth+1)
	sn.ExplodeAll(root.RightNum, depth+1)
}

// Explode explodes the reciever.
func (sn *SnailfishNum) Explode() error {
	if sn.IsLeaf {
		return fmt.Errorf("cannot explode leaf node")
	}
	if !sn.LeftNum.IsLeaf && !sn.RightNum.IsLeaf {
		return fmt.Errorf("cannot explode node with non-leaf children")
	}

	if l := sn.LeftLeafNeighbor(); l != nil {
		l.Value += sn.LeftNum.Value
	}
	if r := sn.RightLeafNeighbor(); r != nil {
		r.Value += sn.RightNum.Value
	}

	zeroNum := &SnailfishNum{
		Value:  0,
		IsLeaf: true,
	}
	parent := sn.Parent
	if parent.LeftNum == sn {
		parent.LeftNum = zeroNum
		parent.LeftNum.Parent = parent
	} else if parent.RightNum == sn {
		parent.RightNum = zeroNum
		parent.RightNum.Parent = parent
	}

	return nil
}

// SplitOne finds one item to split and returns
func (sn *SnailfishNum) SplitOne(root *SnailfishNum) bool {
	if root == nil {
		return false
	}
	if root.IsLeaf && root.Value >= 10 {
		if err := root.Split(); err != nil {
			log.Fatal(err)
		}
		return true
	}
	splitLeft := sn.SplitOne(root.LeftNum)
	if splitLeft {
		return true
	}
	return sn.SplitOne(root.RightNum)
}

// Split performs a split operation on the receiver.
func (sn *SnailfishNum) Split() error {
	if !sn.IsLeaf {
		return fmt.Errorf("cannot split number that isn't a leaf")
	}
	if sn.Value < 10 {
		return fmt.Errorf("cannot split number with value < 10, but value = %d", sn.Value)
	}

	parent := sn.Parent
	splitNum := &SnailfishNum{
		Value: RootVal,
		LeftNum: &SnailfishNum{
			Value:  sn.Value / 2,
			IsLeaf: true,
		},
		RightNum: &SnailfishNum{
			Value:  int(math.Ceil(float64(sn.Value) / 2.0)),
			IsLeaf: true,
		},
		Parent: parent,
	}
	splitNum.LeftNum.Parent = splitNum
	splitNum.RightNum.Parent = splitNum

	if parent.LeftNum == sn {
		parent.LeftNum = splitNum
	} else {
		parent.RightNum = splitNum
	}

	return nil
}

// LeftLeafNeighbor returns the leaf node directly to the left if there is one.
func (sn *SnailfishNum) LeftLeafNeighbor() *SnailfishNum {
	if sn == nil || sn.Parent == nil {
		return nil
	}
	parent := sn.Parent
	if parent.LeftNum != sn {
		if parent.LeftNum.IsLeaf {
			return parent.LeftNum
		} else {
			return furthestRightLeaf(parent.LeftNum)
		}
	}
	return parent.LeftLeafNeighbor()
}

// RightLeafNeighbor returns the leaf node directly to the left if there is one.
func (sn *SnailfishNum) RightLeafNeighbor() *SnailfishNum {
	if sn == nil || sn.Parent == nil {
		return nil
	}
	parent := sn.Parent
	if parent.RightNum != sn {
		if parent.RightNum.IsLeaf {
			return parent.RightNum
		} else {
			return furthestLeftLeaf(parent.RightNum)
		}
	}
	return parent.RightLeafNeighbor()
}

func furthestLeftLeaf(root *SnailfishNum) *SnailfishNum {
	if root == nil {
		return nil
	}
	if root.IsLeaf {
		return root
	}
	return furthestLeftLeaf(root.LeftNum)
}

func furthestRightLeaf(root *SnailfishNum) *SnailfishNum {
	if root == nil {
		return nil
	}
	if root.IsLeaf {
		return root
	}
	return furthestRightLeaf(root.RightNum)
}

// NewSnailfishNum creates a binary tree representation of a snailfish number from a string.
func NewSnailfishNum(s string) *SnailfishNum {
	bracketStack := newStack()
	for i, char := range s {
		switch char {
		case '[':
			bracketStack.Push(char)
		case ']':
			bracketStack.Pop()
		case ',':
			// We are at the root of the current number
			if bracketStack.Len() == 1 {
				var left, right *SnailfishNum
				if isDigit(s[i-1]) {
					var value int
					value, _ = strconv.Atoi(string(s[i-1]))
					left = &SnailfishNum{
						Value:  value,
						IsLeaf: true,
					}
				} else {
					leftStr := getNestedLeft(s[:i])
					left = NewSnailfishNum(leftStr)
				}
				if isDigit(s[i+1]) {
					var value int
					value, _ = strconv.Atoi(string(s[i+1]))
					right = &SnailfishNum{
						Value:  value,
						IsLeaf: true,
					}
				} else {
					rightStr := getNestedRight(s[i+1:])
					right = NewSnailfishNum(rightStr)
				}
				root := &SnailfishNum{
					Value:    RootVal,
					LeftNum:  left,
					RightNum: right,
				}
				left.Parent = root
				right.Parent = root
				return root
			}
		}
	}
	return nil
}

func getNestedLeft(s string) string {
	stackSize := 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case ']':
			stackSize++
		case '[':
			stackSize--
			if stackSize == 0 {
				return s[i:]
			}
		}
	}
	return ""
}

func getNestedRight(s string) string {
	stackSize := 0
	for i, c := range s {
		switch c {
		case '[':
			stackSize++
		case ']':
			stackSize--
			if stackSize == 0 {
				return s[:i+1]
			}
		}
	}
	return ""
}

func (sn *SnailfishNum) String() string {
	return inorderString(sn)
}

func inorderString(root *SnailfishNum) string {
	if root.IsLeaf {
		return fmt.Sprintf("%d", root.Value)
	}
	var fromRoot string
	if root != nil {
		fromRoot += "["
		fromRoot += inorderString(root.LeftNum)
		fromRoot += ","
		fromRoot += inorderString(root.RightNum)
		fromRoot += "]"
	}
	return fromRoot
}

func inorder(root *SnailfishNum) []int {
	var fromRoot []int
	if root != nil {
		fromRoot = append(fromRoot, inorder(root.LeftNum)...)
		fromRoot = append(fromRoot, root.Value)
		fromRoot = append(fromRoot, inorder(root.RightNum)...)
	}
	return fromRoot
}

func isDigit(r byte) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
