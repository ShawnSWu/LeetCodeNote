package main

/*

經過用紙筆試著假設某個階梯數，然後把每種可能窮舉出來
e.g. 可以自己窮舉看看 當:
stair=3時 有3種方法
stair=4時 有5種方法
stair=5時 有8種方法
stair=6時 有13種方法

觀察上方的窮舉結果會發現規律是
「爬到第 n 階樓梯的方法數」會等於「爬上 n−1 階樓梯的方法數量」 + 「爬上 n−2 階樓梯的方法數量」

這類型的公式其實就是『費波那契數列』，公式：
f(n) = f(n-1) + f(n-2)
*/

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	arr := make([]int, n)

	arr[0] = 1
	arr[1] = 2

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n-1]

}

func main() {
	climbStairs(5)
}
