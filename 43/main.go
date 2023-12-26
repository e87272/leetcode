package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Printf("%+v", multiply("100000", "456"))
}

func multiply(num1 string, num2 string) string {
	result := "0"
	if num1 == "0" || num2 == "0" {
		return result
	}

	if len(num1) > len(num2) {
		temp := num1
		num1 = num2
		num2 = temp
	}

	for num1Index := len(num1) - 1; num1Index >= 0; num1Index-- {
		num1Shift := len(num1) - num1Index - 1
		num1Value, _ := strconv.Atoi(string(num1[num1Index]))
		for num2Index := len(num2) - 1; num2Index >= 0; num2Index-- {
			num2Shift := len(num2) - num2Index - 1
			num2Value, _ := strconv.Atoi(string(num2[num2Index]))
			value := num1Value * num2Value
			shiftDigits := num1Shift + num2Shift
			var sum int
			if len(result)-1 < shiftDigits {
				result = strings.Repeat("0", shiftDigits-len(result)+1) + result
			}
			resultIndex := len(result) - shiftDigits - 1
			digitNum, _ := strconv.Atoi(string(result[resultIndex]))
			sum = value + digitNum
			resultBytes := []byte(result)
			resultBytes[resultIndex] = byte(sum%10) + '0'
			result = string(resultBytes)
			for sum/10 > 0 {
				shiftDigits = shiftDigits + 1
				sum = sum / 10
				if len(result) > shiftDigits {
					resultIndex := len(result) - shiftDigits - 1
					digitNum, _ := strconv.Atoi(string(result[resultIndex]))
					sum = sum + digitNum
					resultBytes := []byte(result)
					resultBytes[resultIndex] = byte(sum%10) + '0'
					result = string(resultBytes)
				} else {
					result = string(append([]byte(strconv.Itoa(sum%10)), []byte(result)...))
					sum = 0
				}
			}
		}
	}

	return result
}
