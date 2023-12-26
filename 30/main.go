package main

import "log"

func main() {

	log.Printf("findSubstring : %+v", findSubstring("barfoobarthefoobarman", []string{"foo", "bar"}))
}

func findSubstring(s string, words []string) []int {

	result := []int{}

	copyWords := make([]string, len(words))
	copy(copyWords, words)
	sumLen := 0
	for i := 0; i < len(copyWords); i++ {
		sumLen = sumLen + len(copyWords[i])
	}

	for k := 0; k < len(s); k++ {
		nowIndex := k
		copyWords := make([]string, len(words))
		copy(copyWords, words)
		if len(s[k:]) < sumLen {
			break
		}
		wordOrder := []string{}
		for i := 0; i < len(copyWords); i++ {
			word := copyWords[i]
			if nowIndex+len(word) > len(s) {
				break
			}
			// log.Printf("nowIndex : %+v subStr : %+v ", nowIndex, string(s[nowIndex:nowIndex+len(word)]))
			if s[nowIndex:nowIndex+len(word)] == word {
				wordOrder = append(wordOrder, word)
				copyWords = append(copyWords[:i], copyWords[i+1:]...)
				i = -1
				nowIndex = nowIndex + len(word)
				// log.Printf("len(word) : %+v", len(word))
			}
			// log.Printf("k : %+v copyWords : %+v ", k, copyWords)
		}
		if len(copyWords) == 0 {
			result = append(result, k)
			for _, v := range wordOrder {
				if k+sumLen+len(v) > len(s) || len(v) > 1 {
					break
				}
				// log.Printf("s[k+sumLen+1:k+sumLen+len(v)] : %+v ", string(s[k+sumLen:k+sumLen+len(v)]))
				if s[k+sumLen:k+sumLen+len(v)] == v {
					result = append(result, k+len(v))
					k = k + len(v)
				} else {
					break
				}
			}
		}
	}
	return result
}
