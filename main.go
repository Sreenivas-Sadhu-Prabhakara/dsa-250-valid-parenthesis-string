package main

import "fmt"

// CheckValidString reports whether a string of '(', ')' and '*' can be read
// as balanced, with each '*' acting as '(', ')' or nothing. Track the RANGE
// of possible open-bracket counts instead of exploring choices.
func CheckValidString(s string) bool {
	lo, hi := 0, 0 // fewest and most unmatched '(' possible so far
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			lo++
			hi++
		case ')':
			lo--
			hi--
		default: // '*': could be '(', ')' or empty
			lo--
			hi++
		}
		if hi < 0 {
			return false // too many ')' even with every '*' as '('
		}
		if lo < 0 {
			lo = 0 // can't have fewer than zero opens; clamp the range
		}
	}
	return lo == 0 // some choice closes everything exactly
}

func main() {
	fmt.Println(CheckValidString("(*))"))
	fmt.Println(CheckValidString(")("))
}
