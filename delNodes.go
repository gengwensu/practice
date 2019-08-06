/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	dh := map[int]bool{}
	for _, v := range to_delete {
		dh[v] = true
	}
	ph := map[*TreeNode]*TreeNode{} //key tn, value parent tn
	ans := []*TreeNode{}
	queue := []*TreeNode{root}

	if !dh[root.Val] {
		ans = append(ans, root)
	}

	for len(queue) > 0 {
		next := []*TreeNode{}
		for _, tn := range queue {
			left, right := tn.Left, tn.Right
			if left != nil {
				next = append(next, left)
				ph[left] = tn
			}
			if right != nil {
				next = append(next, right)
				ph[right] = tn
			}
			if dh[tn.Val] { //delete node
				if tn == root {
					root.Left = nil
					root.Right = nil
				} else { //sever children
					tn.Left = nil
					tn.Right = nil
					parent := ph[tn]
					if parent.Left == tn { //sever parent
						parent.Left = nil
					} else {
						parent.Right = nil
					}
				}
				if left != nil && !dh[left.Val] {
					ans = append(ans, left)
				}
				if right != nil && !dh[right.Val] {
					ans = append(ans, right)
				}
			}
			queue = next
		}
	}
	return ans
}
