package main

import (
	"log"
)

func main() {
	s := "cbabca"
	output := longestPalindrome(s)

	log.Printf("output : %+v", output)
}

func longestPalindrome(s string) string {

	startIndex := 0
	endIndex := 1

	for j := 0; j < len(s); j++ {
		if j+1 < len(s) && s[j] == s[j+1] {
			i := 0
			for j+1+i < len(s) && j-i >= 0 && s[j-i] == s[j+1+i] {
				i++
			}
			i = i - 1
			if endIndex-startIndex < len(s[j-i:j+1+i+1]) {
				startIndex = j - i
				endIndex = j + 1 + i + 1
			}
		}

		if j+1 < len(s) && j-1 >= 0 && s[j+1] == s[j-1] {
			i := 0
			for j+1+i < len(s) && j-1-i >= 0 && s[j-1-i] == s[j+1+i] {
				i++
			}
			i = i - 1
			if endIndex-startIndex < len(s[j-1-i:j+1+i+1]) {
				startIndex = j - 1 - i
				endIndex = j + 1 + i + 1
			}
		}

	}

	return s[startIndex:endIndex]

}
