package main

import (
	"fmt"
	"time"
)

func sieveOfEratos(maxSize int, primeReq int) (int, error) {
	isPrime := make([]bool, maxSize) //!! negative logic
	var primeCount []int             // stores prime number less than maxSize
	factor := make([]int, maxSize)   // smallest factor of sieveOfEratos

	isPrime[0], isPrime[1] = !false, !false // 0, 1 is not prime number

	for i := 2; i < maxSize; i++ {
		if !isPrime[i] {
			primeCount = append(primeCount, i)
			factor[i] = i
		}

		if len(primeCount) == primeReq {
			//fmt.Printf("Output: %d\n", i) // use for debug
			return i, nil
		}

		for j := 0; j < len(primeCount) && i*primeCount[j] < maxSize && primeCount[j] <= factor[i]; j++ {
			isPrime[i*primeCount[j]] = !false
			factor[i*primeCount[j]] = primeCount[j]
		}
	}

	return -1, fmt.Errorf("Out of range maxSize\n")
}

func main() {
	var input int

	fmt.Print("Input : ")
	fmt.Scan(&input)
	if input < 1 {
		panic("Input Must >= 1")
	}

	strartTime := time.Now() // start time

	primePn, err := sieveOfEratos(15485868, input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output: %d\n", primePn)

	endTime := time.Now() // end time
	executionTime := endTime.Sub(strartTime)
	fmt.Printf("Function PrimeNumver N position executed in %v\n", executionTime)
}
