package main

/**
Longest Common Prefix
https://leetcode.com/problems/longest-common-prefix/

設一個Point值，當作當前檢查的指針
迴圈走訪每個字串，每輪都透過point取出一個Char比對，若某輪的Point位置在每個字串中都相同，則進到下一輪
但只要每輪中有任何一個Char與其他人不同 或 為空，則需把point-1(代表此輪的Point並非共同擁有)，接著結束迴圈。
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == ""{
		return ""
	}
	//指針
	currentPoint := 0
	//先抓第一陣列第一個當作初始值
	temp := strs[0][0]

	for true {
		//因為是以第一個作為基準 每輪開始前要先確保第一個的Point指到的值不為null
		if len(strs[0]) == currentPoint {
			goto End
		}
		temp = strs[0][currentPoint]

		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == currentPoint || temp != strs[i][currentPoint] {
				goto End
			}
		}
		currentPoint += 1
	}

End:
	currentPoint -= 1

	return strs[0][0:currentPoint+1]

}

func main() {
	var strs = []string{"flower","flow","flight"}
	longestCommonPrefix(strs)
}

