package main

import "math"

func main() {

}

func divide(dividend int, divisor int) int {
	if int(dividend/divisor) > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(dividend / divisor)
}
