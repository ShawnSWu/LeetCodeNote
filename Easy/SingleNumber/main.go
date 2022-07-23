package main

/*
透過Map記錄出現次數即可解決
*/
func singleNumber(nums []int) int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}

	var answer int
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			answer = nums[i]
		}
	}

	return answer
}
