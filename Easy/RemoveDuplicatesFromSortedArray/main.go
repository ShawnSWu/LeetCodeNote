package main

/*
Remove Duplicates from Sorted Array
https://leetcode.com/problems/remove-duplicates-from-sorted-array/

使用兩個指針(Demo中的point與for迴圈的i) 同時指向0，i負責走訪每一個元素，point負責指向『當前最後一個不重複的index』
for的i會自動增加， 每輪都去比較nums[point]與nums[i]是否相同
若不同 -> point+=1 代表都沒遇到重複 即往前進，
若相同 -> 代表遇到重複，此時point不動，再透過一個布林值紀律遇到重複，
再接續跑for迴圈 直到找到不重複的值，再此沒重複的值指向point+1的位置。
 */


func removeDuplicates(nums []int) int {

	point := 0
	isDu := false

	for i:=0; i<len(nums); i++ {

		if nums[i] != nums[point] {
			if isDu {
				point++
				nums[point] = nums[i]
				isDu = false
				i--
			}else {
				point++
			}
		}else {
			isDu = true
		}
	}
	return point + 1
}


func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
}
