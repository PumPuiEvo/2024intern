package main

import (
	"fmt"
	"math"
	"time"
)

var countOut int

func isPrimePositive(number int) bool {
	if number < 2 {
		return false
	}

	if number == 3 || number == 2 {
		return true
	}

	if number%2 == 0 || number%3 == 0 {
		return false
	}

	maxCalculate := int(math.Ceil(math.Sqrt(float64(number))))

	for i := 5; i <= maxCalculate; i += 6 {
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
	if input < 1 {
		panic("Input Must >= 1")
	}

	strartTime := time.Now() // start time

	count := 1
	number := 1
	if input == 1 { // corner case
		number = 2
	}

	for count < input {
		number += 2
		if isPrimePositive(number) {
			count++
		}
	}
	fmt.Printf("Output: %d\n", number)

	endTime := time.Now() // end time
	executionTime := endTime.Sub(strartTime)
	fmt.Printf("Function PrimeNumver N position executed in %v\n", executionTime)
}
