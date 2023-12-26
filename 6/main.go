package main

import "log"

func main() {
	str := convert("PAYPALISHIRING", 3)

	log.Printf("str : %+v", str)
}

func convert(s string, numRows int) string {
	mod := numRows
	if numRows > 2 {
		mod = numRows + (numRows - 2)
	}

	ansStr := make([]string, numRows)
	for i := 0; i < int(len(s)/mod); i++ {
		for k, s := range s[i*mod : i*mod+mod] {
			if k < numRows {
				ansStr[k] = ansStr[k] + string(s)

			} else {
				ansStr[(numRows-2)-k%numRows] = ansStr[(numRows-2)-k%numRows] + string(s)
			}
		}
	}
	for k, s := range s[int(len(s)/mod)*mod:] {
		if k < numRows {
			ansStr[k] = ansStr[k] + string(s)
		} else {
			ansStr[(numRows-2)-k%numRows] = ansStr[(numRows-2)-k%numRows] + string(s)
		}
	}

	// 0 1 2 3 4
	//   7 6 5

	log.Printf("ansStr : %+v", ansStr)
	s = ""
	for _, v := range ansStr {
		s = s + v
	}
	return s
}
