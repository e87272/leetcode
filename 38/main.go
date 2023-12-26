package main

import (
	"log"
	"strconv"
)

func main() {

	log.Printf("%+v", countAndSay(5))

}
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	say := countAndSay(n - 1)
	start := 0
	i := 1
	newSay := ""
	for i < len(say) {
		if say[i] != say[i-1] {
			newSay = newSay + strconv.Itoa(i-start) + strconv.Itoa(int(say[i-1]-'0'))
			start = i
		}
		i++
	}
	newSay = newSay + strconv.Itoa(i-start) + strconv.Itoa(int(say[i-1]-'0'))
	return newSay
}
