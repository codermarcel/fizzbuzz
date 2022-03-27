package main

import "fmt"

func fizzbuzz(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "Fizzbuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", i)
}

func main() {
	for i := 1; i <= 25; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
