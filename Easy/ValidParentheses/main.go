package main

/*
Valid Parentheses
https://leetcode.com/problems/valid-parentheses/

字串從左掃到右，再透過Stack的特性去解，
Stack只存左邊符號，當右邊符號出現時再從Stack取出，看看是否是對應的符號即可。
 */

func isValid(s string) bool {
	pmap := map[rune]int{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := New()

	for i := 0; i < len(s); i++ {
		currentChar := rune(s[i])

		if currentChar == '(' || currentChar == '[' || currentChar == '{' {
			stack.Push(currentChar)
		} else {

			if rune(pmap[currentChar]) != stack.Pop() {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func main (){
	isValid("{({[]})}")
}



type Stack struct {
	list []interface{}
}

func New() *Stack {
	s := new(Stack)
	s.list = make([]interface{}, 0, 8)
	return s
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return tmp
}

func (s *Stack) Push(element interface{}) {
	s.list = append(s.list, element)
}
