package main

import "fmt"

func main() {

	fibonacciSeries := [20]int{}

	fibonacciSeries[0] = 1
	fibonacciSeries[1] = 1

	for index := 2; index < 20; index++ {
		fibonacciSeries[index] = fibonacciSeries[index-2] + fibonacciSeries[index-1]
		fmt.Printf("value in series is %v\n", fibonacciSeries[index])
	}
}
