package main

import "fmt"

// CountingNumber sums numbers from 1 to n with a step of 0.5.
func CountingNumber(n int) float64 {
	var sum float64
	for i := 1.0; i <= float64(n); i += 0.5 {
		sum += i
	}
	return sum
}

//debug
func main() {
	fmt.Println(CountingNumber(10))  // Expected output: 104.5
	fmt.Println(CountingNumber(100)) // Expected output: 10049.5
}
