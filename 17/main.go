package main

import "log"

func main() {
	log.Printf("%+v", letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	var numMap = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	ary := []string{}
	for _, v := range digits {
		str := string(v)
		nowAry := []string{}
		if len(ary) == 0 {
			nowAry = numMap[str]
		} else {
			for _, nowStr := range ary {
				for _, addStr := range numMap[str] {
					nowAry = append(nowAry, nowStr+addStr)
				}
			}
		}
		ary = nowAry
	}
	return ary
}
