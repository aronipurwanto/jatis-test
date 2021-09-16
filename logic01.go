package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Soalt logic 01")

	primeNumber := make(map[int]int)
	index := 0
	for i:=2; i <= 11; i++{
		isPrime := true
		for j:= 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false;
				break;
			}
		}

		if isPrime {
			primeNumber[index] = i
			index++
		}
	}

	fmt.Println(primeNumber)
	array2D := make([][]int, 5)

	for i := range array2D{
		array2D[i] = make([]int, 5)
	}

	for i:=0; i < len(array2D); i++ {
		for j:=0; j < len(array2D[i]); j++ {
			if i >= j{
				fmt.Print(primeNumber[j],"\t")
			}
		}
		fmt.Println()
	}
}