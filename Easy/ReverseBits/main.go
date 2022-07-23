package main

import (
	"fmt"
	"strings"
)

/*
提示：『num & 1 可以取出當前的二進制數(num) 最右邊那個數』。

此題只需要把要翻轉的數從右向左一位位的取出來，如果取出來的是1，將結果(res)左移一位並且加上1，
如果取出來的是0，將結果 res 左移一位，然後將n右移一位即可。

 心得：此題是位移運算，關鍵在弄懂 << >> 的位移運算 與 num & 1 的操作!
*/

func reverseBits(num uint32) uint32 {
	var res uint32 = 0

	//因為只有32位，所以只走訪32次就行了
	for i := 0; i < 32; i++ {

		//透過(num & 1)判斷當前二進制數最右邊那個數
		if (num & 1) == 1 { //若是1
			res = (res << 1) + 1 //則在res向左位移一位，並且+1，(位移本身只會補0，但因為此處要是1，所以才+1)

			//printBits(res) debug用 可以查看每個階段的數值
			//fmt.Println()
		} else { //若不是零
			res = res << 1 //單純位移就好，因為位移本身就會補0

			//printBits(res) debug用 可以查看每個階段的數值
			//fmt.Println()
		}
		num = num >> 1 //右移 把輸入的input右移，代表會在最左邊補0，最後右邊的數就會消失，變成原本的倒數第二個數
		//  e.g. 10010 右移後變成 01001
	}

	return res
}

func main() {

	var num uint32 = 43261596

	reverseBits(num)
}

func printBits(x uint32) {
	var slice []string
	for i := 31; i >= 0; i-- {
		if (x & (1 << i)) != 0 {
			slice = append(slice, "1")
		} else {
			slice = append(slice, "0")
		}
	}
	fmt.Println(strings.Join(slice, ""))
}
