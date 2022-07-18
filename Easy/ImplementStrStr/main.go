package main

/*
將needle當作window，透過此window掃描haystack，直到haystack最後
如果有匹配則return當下index
*/

func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack){
		return -1
	}

	i:=0

	for true {
		window := len(needle)

		if haystack[i:window+i] == needle{
			return i
		}

		if i == (len(haystack) - window) {
			break
		}
		i++
	}

	return -1
}

func main(){
	strStr("hello", "ll")
}
