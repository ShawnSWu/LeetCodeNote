package main

/*
設一組map，在每次計算前都先查表，看看要算的數之前有沒有算過，若沒有則接著計算，
計算結束後存下計算結果，這樣當再度遇到同一個數的時候，就能知道此數絕不可能是happy number.
*/

func isHappy(n int) bool {

	theMap := make(map[int]int)

	for true {

		if theMap[n] == 0 {
			answer := calculate(n)
			theMap[n] = answer

			if answer == 1 {
				return true
			}

			n = answer
		} else {
			break
		}
	}
	return false
}

func calculate(n int) int {
	answer := 0

	for n > 0 {
		temp := n % 10
		answer += pow(temp)
		n = n / 10
	}
	return answer
}

func pow(n int) int {
	return n * n
}

func main() {
	isHappy(7)
}
