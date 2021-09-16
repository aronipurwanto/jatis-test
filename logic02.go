package main

import "fmt"

func main() {
	result :=1
	for i:=1; i<=6; i++ {
		result = result * i
		fmt.Print(result, "\t")
	}
}
