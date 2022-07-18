package main

/*
已知root的情況下，同時把root.left, root.right傳進func，
判斷左節點與右節點是否相同，func中需判斷條件：

1.兩個節點都必須不是null，否則false
2.兩個節點值必須相同，否則false
3.往下檢查時，將節點的各自對稱的left right帶入func

3.若兩個節點都null時，則代表檢查結束，且因為前面一連串的檢查都沒檢查到false，故判斷成對稱，所以return true
 */

func isSymmetric(root *TreeNode) bool {
	return findSymmetric(root.Left, root.Right)
}

func findSymmetric(left *TreeNode, right *TreeNode) bool {
	//只要其中一個為null那就代表一定不是對稱的
	if (left != nil && right == nil) ||  (left == nil && right != nil)  {
		return false
	}

	//代表檢查到最後一個了 但是前面都沒處發false 代表前面都是對稱的
	if left == nil && right == nil {
		return true
	}

	//當兩個節點值都相同時才繼續更深層的節點檢查
	if left.Val == right.Val {
		//在此func中 把left與right，當成第二層的2，即可理解此處程式碼意義。
		/*
		    1
		   / \
		  2   2
		 / \ / \
	 	3  4 4  3

		*/			//所以此處是   3           3                            4            4
		return findSymmetric(left.Left, right.Right) && findSymmetric(left.Right, right.Left)
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//isSymmetric()
}
