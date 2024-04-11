package main

import (
	"fmt"
	"math"
	"time"
)

func isPrimeM2(number int) bool {
	// this func can check Prime number more than2
	maxCalculate := math.Ceil(math.Sqrt(float64(number)))

	for i := 3; i <= int(maxCalculate); i += 2 {
		if (number % i) == 0 {
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
		fmt.Println("Output: 2")
	} else {
		count := 2
		number := 3
		for count < input {
			number = number + 2
			if isPrimeM2(number) {
				count++
			}
		}
		fmt.Printf("Output: %d", number)
	}

	endTime := time.Now() // end time
	executionTime := endTime.Sub(strartTime)
	fmt.Printf("\nFunction executed in %v\n", executionTime)

}
