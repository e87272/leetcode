package main

import (
	"log"
)

func main() {
	log.Printf("%+v", generateParenthesis(4))
}

func generateParenthesis(n int) []string {

	parenthesis := []string{}
	for i := 1; i <= n; i++ {
		parenthesis = append(parenthesis, addParentheses(repeat("(", i), n-i, i, n)...)
	}

	return parenthesis
}

func addParentheses(s string, left int, right int, n int) []string {
	// log.Printf("s : %+v left : %+v right : %+v", s, left, right)
	parenthesis := []string{}
	for j := 1; j <= right; j++ {
		s = s + ")"
		for i := 1; i <= left; i++ {
			parenthesis = append(parenthesis, addParentheses(s+repeat("(", i), left-i, (right-j)+i, n)...)
		}
		// log.Printf("s : %+v left : %+v right : %+v j : %+v", s, left, right, j)
		if len(s) == n*2 {
			parenthesis = append(parenthesis, s)
		}
	}
	return parenthesis
}

func repeat(repeatStr string, n int) string {
	newStr := ""
	for i := 0; i < n; i++ {
		newStr = newStr + repeatStr
	}
	return newStr
}
