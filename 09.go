package main

import "fmt"

func fibonacci(n int) map[int]int {

	fibo := make(map[int]int)

	for i := 0; len(fibo) != n; i++ {

		if i < 2 {

			fibo[i] = i

		} else {

			fibo[i] = fibo[i-2] + fibo[i-1]

		}

	}

	return fibo

}

func main() {

	fmt.Println(fibonacci(20))

}
