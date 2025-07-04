package main

func main() {
	var s string = "{([])}[]"
	isValid(s)
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// var list = make([]rune, len(s))
	var list = []byte{}
	list = append(list, s[0])
	for i := 1; i < len(s); i++ {
		// fmt.Println(string(s[i]))
		if len(list) > 0 && (list[len(list)-1] == pairs[s[i]]) {
			list = list[:len(list)-1]
		} else {
			list = append(list, s[i])
		}
	}
	return len(list) == 0
}
