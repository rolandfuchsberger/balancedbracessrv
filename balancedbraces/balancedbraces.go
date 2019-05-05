package balancedbraces

var closingBraces = map[rune]struct{}{')': struct{}{}, '}': struct{}{}, ']': struct{}{}}

func isClosingBrace(r rune) bool {
	_, ok := closingBraces[r]
	return ok
}

var openingBraces = map[rune]struct{}{'(': struct{}{}, '{': struct{}{}, '[': struct{}{}}

func isOpeningBrace(r rune) bool {
	_, ok := openingBraces[r]
	return ok
}

var bracesPairs = map[rune]rune{
	'(': ')', '[': ']', '{': '}',
}

//BalancedBraces checks for (, ), {, } and [, ]. True if braces are balanced.
func BalancedBraces(s string) bool {

	stack := &stack{}

	for _, e := range s {

		if isOpeningBrace(e) {
			stack.push(e)
		}

		if isClosingBrace(e) && stack.len() == 0 {
			return false
		}

		if isClosingBrace(e) && bracesPairs[stack.pop()] != e { // note: stack.pop() has side effects
			return false
		}
	}

	return stack.len() == 0
}
