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
		for true {
			temp := ((float64(maxUpperbound) / math.Log(float64(maxUpperbound))) / float64(input))
			if temp > 1 {
				break
			}

			if maxUpperbound > math.MaxInt32 {
				return fmt.Errorf("the value input too high")
			}
			maxUpperbound += 100000000 // + 100m until MaxInt32
		}

	} else {
		listOfInput := [][]int{
			{500001, 7368791 + 1},
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

	primePn, err := sieveOfEratos(int(maxUpperbound+1), input)
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
