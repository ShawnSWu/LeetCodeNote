package main

/*
二元搜尋數中序走訪
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var answerArr = []int{}

func inorderTraversal(root *TreeNode) []int {
	if root != nil {
		inorderTraversal(root.Left)
		answerArr = append(answerArr, root.Val)
		inorderTraversal(root.Right)
	}
	return answerArr
}
