package main

import (
	"fmt"
	"math"
	"time"
)

func sieveOfEratos(maxSize int, primeReq int) (int, error) {
	isNotPrime := make([]bool, maxSize)
	var primeCount []int           // stores prime number less than maxSize
	factor := make([]int, maxSize) // smallest factor of sieveOfEratos

	isNotPrime[0], isNotPrime[1] = true, true // 0, 1 is not prime number

	for i := 2; i < maxSize; i++ {
		if !isNotPrime[i] { // is prime number
			primeCount = append(primeCount, i)
			factor[i] = i
		}

		if len(primeCount) == primeReq {
			//fmt.Printf("Output: %d\n", i) // use for debug
			return i, nil
		}

		for j := 0; j < len(primeCount) && i*primeCount[j] < maxSize && primeCount[j] <= factor[i]; j++ {
			isNotPrime[i*primeCount[j]] = true
			factor[i*primeCount[j]] = primeCount[j]
		}
	}

	return -1, fmt.Errorf("Out of range maxSize\n")
}

func inputHandle(input int) error {
	if input < 1 { // corner case
		return fmt.Errorf("Input Must >= 1")
	}

	var maxUpperbound int
	maxPreCase := 10000000
	if input > maxPreCase {
		maxUpperbound = 247483647 // 247,483,647 + 1900M is MaxInt32

		for (maxUpperbound/(int(math.Log10(float64(maxUpperbound))+1)))/input < 3 {
			// ((float64(maxUpperbound) / math.Log(float64(maxUpperbound))) / float64(input)) > 1 true formular
			maxUpperbound += 100000000 // + 100m until MaxInt32
			if maxUpperbound > math.MaxInt32 {
				return fmt.Errorf("the value input too high")
			}
		}

	} else {
		listOfInput := [][]int{
			{1000001, 15485867 + 1},
			{2000001, 32452867 + 1},
			{5000001, 86028157 + 1},
			{maxPreCase, 179424673 + 1}}
		// case for reserve memory

		for i := 0; i < len(listOfInput); i++ {
			if input <= listOfInput[i][0] {
				maxUpperbound = listOfInput[i][1]
				break
			}
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
