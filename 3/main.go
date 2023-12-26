package main

import (
	"log"
	"strings"
)

func main() {
	s := "pwwkew"
	output := lengthOfLongestSubstring(s)

	log.Printf("%+v", output)
}

func lengthOfLongestSubstring(s string) int {
	ary := strings.Split(s, "")
	temp := []string{}
	output := 0
	repeatFlag := false

	for _, v1 := range ary {
		repeatFlag = false
		repeatIndex := 0
		for k, v2 := range temp {
			if v2 == v1 {
				repeatFlag = true
				repeatIndex = k + 1
				break
			}
		}
		// log.Printf("repeatIndex : %+v", repeatIndex)
		if repeatFlag {
			if len(temp) > output {
				output = len(temp)
				// log.Printf("temp : %+v", temp)
			}
			temp = temp[repeatIndex:]
			temp = append(temp, v1)
			// log.Printf("temp : %+v", temp)
		} else {
			temp = append(temp, v1)
		}
	}
	if !repeatFlag {
		if len(temp) > output {
			output = len(temp)
			// log.Printf("temp : %+v", temp)
		}
	}

	return output
}
