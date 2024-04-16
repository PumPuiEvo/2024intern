package main

import (
	"fmt"
	"math"
	"time"
)

func sieveOfAtkin(limit int, input int) (int, error) {
	sieve := make([]bool, limit)

	if limit > 2 {
		sieve[2] = true
	}
	if limit > 3 {
		sieve[3] = true
	}

	for x := 1; x*x <= limit; x++ {
		for y := 1; y*y <= limit; y++ {

			// Condition 1
			var n int = (4 * x * x) + (y * y)
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				// sieve[n] = ~sieve[n]
				if sieve[n] {
					sieve[n] = false
				} else {
					sieve[n] = true
				}
			}
			// Condition 2
			n = (3 * x * x) + (y * y)
			if n <= limit && n%12 == 7 {
				// sieve[n] = ~sieve[n]
				if sieve[n] {
					sieve[n] = false
				} else {
					sieve[n] = true
				}
			}

			// Condition 3
			n = (3 * x * x) - (y * y)
			if x > y && n <= limit && n%12 == 11 {
				// sieve[n] = ~sieve[n]
				if sieve[n] {
					sieve[n] = false
				} else {
					sieve[n] = true
				}
			}

		}
	}

	// Mark all multiples of squares as non-prime
	for r := 5; r*r <= limit; r++ {
		if sieve[r] {
			for i := r * r; i <= limit; i += r * r {
				sieve[i] = false
			}
		}
	}

	countPrime := 0
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			countPrime++
		}
		if countPrime == input {
			return i, nil
		}
	}

	return -1, fmt.Errorf("Out of range maxSize\n")
}

func inputHandle2(input int) error {
	if input < 1 { // corner case
		return fmt.Errorf("Input Must >= 1")
	}

	var maxUpperbound float64 = 100000
	for true {
		// prime number theorem say : lim x-> inf (primeCount(x) / xlnx) = 1 (decresses to 1)
		temp := (maxUpperbound / math.Log(maxUpperbound)) / float64(input)
		if temp > 1 {
			break
		}

		if maxUpperbound > math.MaxInt32 {
			return fmt.Errorf("the value input too high")
		}
		maxUpperbound += 100000 // until MaxInt32
	}

	primePn, err := sieveOfAtkin(int(maxUpperbound+2), input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output: %d\n", primePn)
	return nil
}

func main() {
	var input int

	fmt.Print("Input : ")
	fmt.Scan(&input)

	strartTime := time.Now() // start time

	err := inputHandle2(input)
	if err != nil {
		panic(err)
	}

	endTime := time.Now() // end time
	executionTime := endTime.Sub(strartTime)
	fmt.Printf("Function PrimeNumver N position executed in %v\n", executionTime)
}
