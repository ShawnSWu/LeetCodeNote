package main

import "fmt"

/*
Plus One
https://leetcode.com/problems/plus-one/

用i指針從最後面掃回來，判斷+1後<=9，則此處+1後就結束程式，
若+1後>9，則此處變0，i再繼續往下一位判斷
*/

func plusOne(digits []int) []int {
	lastIndex := len(digits)-1

	for i := lastIndex; i >= 0; i-- {
		if (digits[i]+1) <= 9{
			digits[i] = digits[i]+1
			break
		}else {
			digits[i] = 0

			if i == 0 {
				temp := digits

				newDigits := []int{1}

				for i:=0;i<len(temp);i++{
					newDigits = append(newDigits, temp[i])
				}
				return newDigits
			}
		}
	}

	return digits
}

func main() {

	arr := []int{1,2,3,9}
	fmt.Println(plusOne(arr))

}
