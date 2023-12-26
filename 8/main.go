package main

import (
	"log"
	"strconv"
)

func main() {
	a := myAtoi("1")
	log.Printf("a : %+v", a)

}
func myAtoi(s string) int {

	var start int = 0
	var end int = 0
	var sign string = "+"
	var firstFlag bool = false
	for k, v := range s {
		s := string(v)
		if !firstFlag {
			switch s {
			case " ":
				continue
			case "-":
				firstFlag = true
				start = k
				end = k + 1
				sign = "-"
				continue
			case "+":
				firstFlag = true
				start = k
				end = k + 1
				continue
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				firstFlag = true
				start = k
				end = k + 1
				continue
			default:
				return 0
			}
		} else {
			exitFlag := false
			switch s {
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				end = k + 1
				continue
			default:
				exitFlag = true
			}
			if exitFlag {
				end = k
				break
			}
		}
	}

	if end-start == 1 || end-start == 0 {
		switch s[start:end] {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			break
		default:
			return 0
		}
	}
	num, err := strconv.ParseInt(s[start:end], 10, 32)
	if err != nil {
		if sign == "+" {
			return 2147483647
		}
		return -2147483648
	}
	return int(num)
}
