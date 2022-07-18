package main
/*
Merge Sorted Array
https://leetcode.com/problems/merge-sorted-array/

既然是兩個已經排序過的陣列，就從 nums1 和 nums2陣列的末尾開始一個個比較，
把較大的數，按順序放到答案陣列(從後往前放)，直到某個陣列被走訪完時則代表排序完成，
這裡用到三個指標 pm,pn,index，分別指向nums1,nums2和nums1陣列的末尾

每輪完成時再依判斷將這三個指標減一
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {

	pm := m-1
	pn := n-1
	index := len(nums1)-1

	for true {
		if pm < 0 {
			for i:=pn; i>=0; i-- { //當nums1已經被走訪完的狀況，把剩下的nums2的數都加進去
				nums1[i] = nums2[pn]
				index--
				pn--
			}
			break
		}

		if pn < 0 {
			break
		}


		if nums1[pm] <= nums2[pn] {
			nums1[index] = nums2[pn]
			pn--
			index--
		}else{
			nums1[index] = nums1[pm]
			pm--
			index--
		}
	}
}

func main() {

	nums1 := []int{0,0,3,0,0,0,0,0,0}
	m := 3
	nums2 := []int{-1,1,1,1,2,3}
	n := 6
	merge(nums1, m, nums2, n)
}