package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaabbbb"
	fmt.Println("test 01")
	splitString(s)

	s="xyzsaab"
	fmt.Println("\ntest 02")
	splitString(s)
}

func splitString(s string)  {
	slices := strings.Split(s, "")
	result := make(map[string]int)

	for _, slice := range slices{
		if result[slice] == 0 {
			result[slice] = 1
		}else {
			v := result[slice] +1
			result[slice] = v
		}
	}

	for key, item := range result {
		fmt.Print(key, "=", item, "\t")
	}
}