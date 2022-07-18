package main


/*
可以用Divide and conquer概念，
對root節點來說 Depth = findDepth(subTree) + 1
解題思路是 『每個子樹的深度總和 = 此樹的深度』
因此可透過遞迴走訪到最深的子樹，當node == null時代表走到底了，此時回傳深度0。
第二行公式可表示該解法
*/


func maxDepth(root *TreeNode) int {
	return findDepth(root)
}

func findDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftDep := findDepth(node.Left)
	rightDep := findDepth(node.Right)

	if leftDep > rightDep {
		return leftDep
	}else {
		return rightDep
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}