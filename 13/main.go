package main

import "log"

func main() {
	romanInt := romanToInt("MCMXCIV")
	log.Printf("roman : %+v", romanInt)
}

var symbol = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var sum = 0
	var nowRank = 1
	for i := len(s) - 1; i >= 0; i-- {
		v := symbol[string(s[i])]
		if v < nowRank {
			sum = sum - v
		} else {
			sum = sum + v
			nowRank = v
		}
	}
	return sum
}
