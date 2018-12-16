package main

import (
	"fmt"
	"github.com/lm1458777/FizzBuzz/fizzbuzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.NumberToString(i))
	}
}
