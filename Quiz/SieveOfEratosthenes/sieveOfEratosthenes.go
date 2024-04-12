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

func inputHandle(input int) error {
	if input < 1 { // corner case
		return fmt.Errorf("Input Must >= 1")
	}

	if input > 10000000 {
		return fmt.Errorf("the value input too high")
	}

	listOfInput := [][]int{
		{1000001, 15485867 + 1},
		{2000001, 32452867 + 1},
		{5000001, 86028157 + 1},
		{10000000, 179424673 + 1}}
	// case for reserve memory
	var maxUpperbound int
	for i := 0; i < len(listOfInput); i++ {
		if input <= listOfInput[i][0] {
			maxUpperbound = listOfInput[i][1]
			break
		}
	}

	primePn, err := sieveOfEratos(maxUpperbound, input)
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

	err := inputHandle(input)
	if err != nil {
		panic(err)
	}

	endTime := time.Now() // end time
	executionTime := endTime.Sub(strartTime)
	fmt.Printf("Function PrimeNumver N position executed in %v\n", executionTime)
}
