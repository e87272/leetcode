package main

import (
	"log"
)

func main() {
	log.Printf("%+v", longestValidParentheses(")(((((()())()()))()(()))("))
}

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	longest := 0
	count := 0
	firstCount := map[int]int{
		0: -1,
	}
	preCount := 0
	for k, v := range s {
		if v == '(' {
			count = count + 1
		} else {
			count = count - 1
		}
		if k > 0 && count < preCount {
			delete(firstCount, preCount)
		}
		if count < 0 {
			delete(firstCount, 0)
		}
		index, ok := firstCount[count]
		if !ok {
			firstCount[count] = k
		} else {
			if k-index > longest {
				longest = k - index
			}
		}
		preCount = count
	}
	return longest
}
