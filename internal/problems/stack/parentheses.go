package stack

import (
	"fmt"
	"strings"
)

/*
You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:

Every open bracket is closed by the same type of close bracket.
Open brackets are closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Return true if s is a valid string, and false otherwise.

s = ((()))
stack = (((
r = )

*/

var brackets = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isCloseBracket(r rune) bool {
	return brackets[r] != rune(0)
}

func isOpenBracket(r rune) bool {
	return !isCloseBracket(r)
}

func runeStackToString(stack []rune) string {
	var builder strings.Builder
	builder.WriteRune('\'')
	for _, r := range stack {
		builder.WriteRune(r)
	}
	builder.WriteRune('\'')
	return builder.String()
}

func IsValidParentheses(s string) bool {
	stack := []rune{}
	for i, r := range []rune(s) {

		if isOpenBracket(r) {
			stack = append(stack, r)
			fmt.Printf("i=%d, r=%c, stack=%s\n", i, r, runeStackToString(stack))
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != brackets[r] {
				return false
			}
			stack = stack[:len(stack)-1]
			fmt.Printf("i=%d, r=%c, stack=%s\n", i, r, runeStackToString(stack))
		}
	}

	return len(stack) == 0
}
