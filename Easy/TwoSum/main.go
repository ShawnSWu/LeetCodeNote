package main


/*
Two Sum
https://leetcode.com/problems/two-sum/submissions/

1.建立一個Map，Key存值 value存index
2.走訪array，將target-走訪的值= currentNumber
3.在Map中找有無key=currentNumber的資料
4.若有值 則找到解答(原因 -> 簡單的數學加減法)
5.若無 將此 走訪的值 加到map中 繼續2.
*/

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	answerSlice := make([]int, 0, 2)

	for i := 0; i < len(nums); i++ {

		number := target - nums[i]
		index, exist := indexMap[number]

		if !exist {
			indexMap[nums[i]] = i
		} else {
			answerSlice = append(answerSlice, i, index)
		}
	}
	return answerSlice
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
}
