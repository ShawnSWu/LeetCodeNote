package main

/*
找出股票價錢最低的日期與價錢最高的日期即可，但上述條件必須符合『 最低價的日子 必須小於 最高價的日子』，
因為常理來說你無法在第四天買到第三天的價錢，舉例來說 如果第三天10$ 第四天5$，那就不符合條件

實際作法：
透過變數min, max 來分別表示 最低價的日子, 最高價的日，
兩個一起從0出發，當min的價錢 > max的價錢，則min+1，(代表此min無意義 因為大於max，所以往右移動尋找其他可能)
反之 則max+1(代表min暫時找到最小的值 所以讓max往右探索其他可能)
且在過程中紀錄max min的收益 並儲存最大值，
反覆此做法 一直到max走到陣列盡頭，
*/

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {

	minIndex := 0
	maxIndex := 0

	maxProfit := 0

	for maxIndex < len(prices) {

		profit := prices[maxIndex] - prices[minIndex]
		maxProfit = max(maxProfit, profit)

		if prices[minIndex] > prices[maxIndex] {
			minIndex++
		} else {
			maxIndex++
		}
	}
	return maxProfit
}
