package main

import (
	"log"
)

func main() {
	log.Printf("removeDuplicates : %+v", strStr("leetcode", "eetl"))
}
func strStr(haystack string, needle string) int {
	for k, _ := range haystack {
		if k+len(needle) <= len(haystack) && haystack[k:k+len(needle)] == needle {
			return k
		}
	}
	return -1
}
