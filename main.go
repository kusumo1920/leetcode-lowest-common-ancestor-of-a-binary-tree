package main

import "fmt"

func main() {
	q := &TreeNode{
		Val: 4,
	}
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: q,
		},
	}
	input := &TreeNode{
		Val:  3,
		Left: p,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	output := lowestCommonAncestor(input, p, q)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode
	var recursiveFn func(node, p, q *TreeNode) (*TreeNode, *TreeNode)
	recursiveFn = func(node, p, q *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		}

		var pp, qq *TreeNode
		p1, q1 := recursiveFn(node.Left, p, q)
		p2, q2 := recursiveFn(node.Right, p, q)
		if p1 != nil || p2 != nil || node == p {
			pp = p
		}
		if q1 != nil || q2 != nil || node == q {
			qq = q
		}

		if pp != nil && qq != nil && result == nil {
			result = node
		}

		return pp, qq
	}
	recursiveFn(root, p, q)
	return result
}
