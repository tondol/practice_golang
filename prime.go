package main

import (
	"fmt"
)

// n までの素数をエラトステネスのふるいを使って列挙する
func eratosthenes(n int) {
	// 要素数が定数じゃない場合は make を使わないとエラー
	table := make([]bool, n + 1)

	// ゼロ値以外で埋めるにはこの方法しかないのか??
	for i := 2; i < len(table); i++ {
		table[i] = true
	}

	for i := 2; i * i <= n; i++ {
		if !table[i] {
			continue
		}
		for j := 2; i * j <= n; j++ {
			table[i * j] = false
		}
	}

	for i := 1; i <= n; i++ {
		if table[i] {
			fmt.Printf("%d is prime\n", i)
		} else {
			fmt.Printf("%d is not prime\n", i)
		}
	}
}

// n が素数なら true, 合成数なら false を返す
func is_prime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	// sqrt(n) までの奇数で割り算の余りを計算すれば十分
	for i := 3; i * i <= n; i += 2 {
		// fmt.Printf("%d %% %d = %d\n", n, i, n % i)
		if n % i == 0 {
			return false
		}
	}
	return true
}

func is_prime_print(n int) {
	if is_prime(n) {
		fmt.Printf("%d is prime\n", n)
	} else {
		fmt.Printf("%d is not prime\n", n)
	}
}

func main() {
	fmt.Println("-- is_prime")

	is_prime_print(1)
	is_prime_print(2)
	is_prime_print(500)
	is_prime_print(501)
	is_prime_print(3571)

	fmt.Println("-- eratosthenes")

	eratosthenes(100)
}
