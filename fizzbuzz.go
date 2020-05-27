package main

import (
	"fmt"
)

func fizzbuzz_elem(i int) string {
	if i % 15 == 0 {
		return "fizzbuzz!"
	} else if i % 5 == 0 {
		return "buzz"
	} else if i % 3 == 0 {
		return "fizz"
	} else {
		return fmt.Sprintf("%d", i)
	}
}

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(fizzbuzz_elem(i))
	}
}

func main() {
	fizzbuzz(30)
}

// As compiler
// $ go build fizzbuzz.go
// $ ./fizzbuzz

// As interpreter
// $ go run fizzbuzz.go
