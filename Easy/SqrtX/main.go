package main
/*
此題意義上只能一個數一個數去做平方後 判斷是否為此平方根，

但若是從1,2,3,4...這樣一個個看 效率太差，可以透過二分搜尋法 將速度提升
e.g. mySqrt(x=100)
可以先找出二分搜尋的mid(50)
接著判斷50*50是否為x，若不是 則繼續二分(25)
直到right-left == 1時結束，因為如果==1時代表已逼近解答
*/
func mySqrt(x int) int {

	if x==0 {
		return 0
	}
	if x==1 {
		return 1
	}

	//設定二分搜尋的上下界線
	left := 0
	right := x

	mid := -1

	for true {
		if right-left == 1 {
			break
		}

		mid = (left+right) / 2 //二分

		if mid*mid > x {
			right = mid
		}else{
			left = mid
		}
	}

	return left
}

func main(){
	mySqrt(100)
}