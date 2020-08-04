package _0_Valid_Parentheses

func isLeft(c uint8) bool {
	return ('{' == c) || ('[' == c) || ('(' == c)
}

func isRight(c uint8) bool {
	return ('}' == c) || (']' == c) || (')' == c)
}

func isOk(c1 uint8, c2 uint8) bool {
	return ('{' == c1 && c2 == '}') || ('[' == c1 && ']' == c2) || ('(' == c1 && ')' == c2)
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]uint8, len(s))
	current := -1
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			current++
			stack[current] = s[i]
		} else if isRight(s[i]) {
			if current != -1 && isOk(stack[current], s[i]) {
				current--
			} else {
				return false
			}
		}
	}
	return current == -1
}
