package main

import (
	"log"
)

func main() {
	log.Printf("%+v", lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {

	startFlag := s[len(s)-1] != ' '
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			startFlag = true
			count++
		} else if startFlag {
			break
		}
	}

	return count
}
