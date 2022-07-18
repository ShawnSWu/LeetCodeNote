package main

/*
Pascal's Triangle
https://leetcode.com/problems/pascals-triangle/

畫圖仔細觀察後可以發現，這種三角形的特徵有：
1.第i層有i個元素
2.每層第一個以及最後一個元素值為1
3.每層裡的第n個數(除了頭跟尾)，都是上一層的 n-1與n 相加而來的，

知道規律後就很簡單了，
公式：
arr[i][n] = arr[i-1][n-1] + arr[i-1][n]
*/

func generate(numRows int) [][]int {
	//初始化一個全都是1的三角形
	triangle := initArray(numRows)

	for i := 1; i < numRows; i++ {
		layer := triangle[i]

		//判斷是否為最後一層
		if i == (numRows - 1) {
			break
		}

		for j := 0; j < len(layer)-1; j++ {
			leftIndex := j
			rightIndex := j + 1

			left := layer[leftIndex]
			right := layer[rightIndex]

			triangle[i+1][rightIndex] = left + right
		}
	}
	return triangle
}

func initArray(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		for j := 1; j <= i; j++ {
			triangle[i-1] = make([]int, j)
		}
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			triangle[i][j] = 1
		}
	}
	return triangle
}

func main() {
	generate(5)
}
