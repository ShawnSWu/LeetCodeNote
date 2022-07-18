package main

/*
透過Divide and Conquer概念解題，
可以先比較兩個List的頭哪個最小，如果list1的值小於list2的值的話，那麼list1應該要當頭，
接下來我們要拿list1.next跟list2來做比較大小，比較小的，就會是新的l1連接的下一個節點，
剛好可以用遞迴的方式。
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//mergeTwoLists(list1, list2)
}
