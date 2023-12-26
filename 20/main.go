package main

import "log"

func main() {
	b := isValid("([)]")
	log.Printf("b : %+v", b)
}

func isValid(s string) bool {
	var strType []string
	for _, v := range s {
		switch string(v) {
		case "(", "[", "{":
			strType = append(strType, string(v))
		case ")":
			if len(strType) == 0 || strType[len(strType)-1] != "(" {
				return false
			}
			strType = strType[:len(strType)-1]
		case "]":
			if len(strType) == 0 || strType[len(strType)-1] != "[" {
				return false
			}
			strType = strType[:len(strType)-1]
		case "}":
			if len(strType) == 0 || strType[len(strType)-1] != "{" {
				return false
			}
			strType = strType[:len(strType)-1]
		}
	}
	if len(strType) != 0 {
		return false
	}
	return true
}
