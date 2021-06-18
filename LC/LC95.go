package LC

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	. "GoAlgorithm/Common"
	"fmt"
)

func generateTrees(n int) []*TreeNode {
	return getRestNumRoot(1, n)
}

// index大小的结点作为父结点时，所得到的子树 返回根节点数组
// l.r 表示可选择的数字范围
func getRestNumRoot(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{}
	}
	if l == r {
		return []*TreeNode{
			&TreeNode{Val: l},
		}
	}
	ret := []*TreeNode{}
	for m := l; m <= r; m++ {

		leftTreeSet := getRestNumRoot(l, m-1)
		rightTreeSet := getRestNumRoot(m+1, r)

		if len(leftTreeSet) == 0 {
			for i := 0; i < len(rightTreeSet); i++ {
				root := &TreeNode{
					Val:   m,
					Left:  nil,
					Right: nil,
				}
				root.Right = rightTreeSet[i]
				ret = append(ret, root)
			}
		} else if len(rightTreeSet) == 0 {
			for i := 0; i < len(leftTreeSet); i++ {
				root := &TreeNode{
					Val:   m,
					Left:  nil,
					Right: nil,
				}
				root.Left = leftTreeSet[i]
				ret = append(ret, root)
			}
		} else {
			for i := 0; i < len(leftTreeSet); i++ {
				for j := 0; j < len(rightTreeSet); j++ {
					root := &TreeNode{
						Val:   m,
						Left:  nil,
						Right: nil,
					}
					root.Left = leftTreeSet[i]
					root.Right = rightTreeSet[j]
					ret = append(ret, root)
				}
			}
		}
	}
	return ret
}

func GenerateTreesDp(n int) []*TreeNode {
	dp := make([][][]*TreeNode, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]*TreeNode, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][i] = []*TreeNode{
			&TreeNode{i, nil, nil},
		}
	}

	for k := 2; k <= n; k++ {
		p := 1
		q := k
		for p <= n && q <= n {
			dp[p][q] = []*TreeNode{}
			for e := p; e <= q; e++ {
				fmt.Println(k, p, q, e)
				var leftTreeSet []*TreeNode
				if e-1 < 1 {
					leftTreeSet = []*TreeNode{}
				} else {
					leftTreeSet = dp[p][e-1]
				}
				var rightTreeSet []*TreeNode
				if e+1 > n {
					rightTreeSet = []*TreeNode{}
				} else {
					rightTreeSet = dp[e+1][q]
				}

				if leftTreeSet == nil || len(leftTreeSet) == 0 {
					for i := 0; i < len(rightTreeSet); i++ {
						root := &TreeNode{
							Val:   e,
							Left:  nil,
							Right: nil,
						}
						root.Right = rightTreeSet[i]
						dp[p][q] = append(dp[p][q], root)
					}
				} else if rightTreeSet == nil || len(rightTreeSet) == 0 {
					for i := 0; i < len(leftTreeSet); i++ {
						root := &TreeNode{
							Val:   e,
							Left:  nil,
							Right: nil,
						}
						root.Left = leftTreeSet[i]
						dp[p][q] = append(dp[p][q], root)
					}
				} else {
					for i := 0; i < len(leftTreeSet); i++ {
						for j := 0; j < len(rightTreeSet); j++ {
							root := &TreeNode{
								Val:   e,
								Left:  nil,
								Right: nil,
							}
							root.Left = leftTreeSet[i]
							root.Right = rightTreeSet[j]
							dp[p][q] = append(dp[p][q], root)
						}
					}
				}
			}
			p++
			q++
		}
	}

	return dp[1][n]
}
