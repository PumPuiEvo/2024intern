package main

import (
	"fmt"
	"math"
	"time"
)

var countOut int

func isPrimeM2(number int) bool {
	// this func can check Prime number more than2
	maxCalculate := math.Ceil(math.Sqrt(float64(number)))

	if number == 3 {
		return true
	}

	if number%2 == 0 || number%3 == 0 {
		return false
	}

	for i := 5; i <= int(maxCalculate); i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	var input int

	fmt.Print("Input : ")
	fmt.Scan(&input)

	strartTime := time.Now() // start time

	if input == 1 {
		fmt.Println("Output: 2") // Corner case
	} else {
		count := 2
		number := 3
		for count < input {
			number += 2
			if isPrimeM2(number) {
				count++
			}
		}
		fmt.Printf("Output: %d", number)
	}

	endTime := time.Now() // end time
	executionTime := endTime.Sub(strartTime)
	fmt.Printf("\nFunction PrimeNumver N position executed in %v\n", executionTime)
}
