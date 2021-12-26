package week3

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	snailfishNum1 = &SnailfishNum{
		Value: 10,
		LeftNum: &SnailfishNum{
			Value: 10,
			LeftNum: &SnailfishNum{
				Value:   10,
				LeftNum: &SnailfishNum{Value: 0},
				RightNum: &SnailfishNum{
					Value:    10,
					LeftNum:  &SnailfishNum{Value: 4},
					RightNum: &SnailfishNum{Value: 5},
				},
			},
			RightNum: &SnailfishNum{
				Value:    10,
				LeftNum:  &SnailfishNum{Value: 0},
				RightNum: &SnailfishNum{Value: 0},
			},
		},
		RightNum: &SnailfishNum{
			Value: 10,
			LeftNum: &SnailfishNum{
				Value: 10,
				LeftNum: &SnailfishNum{
					Value:    10,
					LeftNum:  &SnailfishNum{Value: 4},
					RightNum: &SnailfishNum{Value: 5},
				},
				RightNum: &SnailfishNum{
					Value:    10,
					LeftNum:  &SnailfishNum{Value: 2},
					RightNum: &SnailfishNum{Value: 6},
				},
			},
			RightNum: &SnailfishNum{
				Value:    10,
				LeftNum:  &SnailfishNum{Value: 9},
				RightNum: &SnailfishNum{Value: 5},
			},
		},
	}

	snailfishNum2 = &SnailfishNum{
		Value:   10,
		LeftNum: &SnailfishNum{Value: 7},
		RightNum: &SnailfishNum{
			Value:   10,
			LeftNum: &SnailfishNum{Value: 5},
			RightNum: &SnailfishNum{
				Value: 10,
				LeftNum: &SnailfishNum{
					Value:    10,
					LeftNum:  &SnailfishNum{Value: 3},
					RightNum: &SnailfishNum{Value: 8},
				},
				RightNum: &SnailfishNum{
					Value:    10,
					LeftNum:  &SnailfishNum{Value: 1},
					RightNum: &SnailfishNum{Value: 4},
				},
			},
		},
	}

	snailfishNum3 = &SnailfishNum{
		Value:    10,
		LeftNum:  &SnailfishNum{Value: 2},
		RightNum: &SnailfishNum{Value: 9},
	}
)

func cmpOpts(t *testing.T) []cmp.Option {
	t.Helper()
	toInorder := cmp.Transformer("ToInorder", func(in *SnailfishNum) []int {
		return inorder(in)
	})
	return []cmp.Option{toInorder}
}

func TestNewSnailfishNum(t *testing.T) {
	testCases := []struct {
		input string
		want  *SnailfishNum
	}{
		{
			input: "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			want:  snailfishNum1,
		},
		{
			input: "[7,[5,[[3,8],[1,4]]]]",
			want:  snailfishNum2,
		},
		{
			input: "[2,9]",
			want:  snailfishNum3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := NewSnailfishNum(tc.input)
			if diff := cmp.Diff(tc.want, got, cmpOpts(t)...); diff != "" {
				t.Errorf("NewSnailfishNum(%q) returned an unexpected diff (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}

func TestExplodeAll(t *testing.T) {
	testCases := []struct {
		inputStr string
		want     string
	}{
		{
			inputStr: "[[[[[9,8],1],2],3],4]",
			want:     "[[[[0,9],2],3],4]",
		},
		{
			inputStr: "[7,[6,[5,[4,[3,2]]]]]",
			want:     "[7,[6,[5,[7,0]]]]",
		},
		{
			inputStr: "[[6,[5,[4,[3,2]]]],1]",
			want:     "[[6,[5,[7,0]]],3]",
		},
		{
			inputStr: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			want:     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.inputStr, func(t *testing.T) {
			input := NewSnailfishNum(tc.inputStr)
			input.ExplodeAll(input, 0)
			if got := input.String(); got != tc.want {
				t.Errorf("ExplodeAll(%q) = %q, want = %q", tc.inputStr, got, tc.want)
			}
		})
	}
}

func TestSplitOne(t *testing.T) {
	root := NewSnailfishNum("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
	root.ExplodeAll(root, 0)
	root.SplitOne(root)
	want1 := "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"
	if got := root.String(); got != want1 {
		t.Errorf("SplitOne = %q, want = %q", got, want1)
	}

	root.SplitOne(root)
	want2 := "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"
	if got := root.String(); got != want2 {
		t.Errorf("SplitOne = %q, want = %q", got, want2)
	}
}
