package main

import "log"

func main() {
	v := myPow(8.95371, -1)
	log.Printf("%+v", v)
}

func myPow(x float64, n int) float64 {

	var sum float64
	if n > 0 {
		sum = multi(x, n)
	} else if n < 0 {
		n = -n
		sum = multi(x, n)
		sum = 1 / sum
	} else {
		return 1
	}
	return sum
}

func multi(x float64, n int) float64 {
	sum := x
	v, d := m(n, 0)
	// log.Printf("n: %+v v : %+v d : %+v", n, v, d)
	for i := 0; i < d; i++ {
		sum = sum * sum
	}
	// log.Printf("sum : %+v", sum)
	if n-v >= 2 {
		sum = sum * multi(x, n-v)
	} else if n-v == 1 {
		sum = sum * x
	}
	return sum
}

func m(n int, t int) (int, int) {
	if n < 2 {
		return 1, t
	}
	v, d := m(n/2, t+1)
	return 2 * v, d
}
